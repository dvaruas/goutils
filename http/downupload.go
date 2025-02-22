package httputils

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"os"
)

func DownloadURLToPath(
	ctx context.Context,
	filePath string,
	downloadURL string,
) error {
	out, err := os.Create(filePath)
	if err != nil {
		return fmt.Errorf("DownloadURLToPath: %w", err)
	}
	defer out.Close()

	statusCode, respBody, err := DoHTTPStreamedCommunication(
		ctx,
		downloadURL,
		http.MethodGet,
		http.NoBody,
		nil)
	if err != nil {
		return fmt.Errorf("DownloadURLToPath: %w", err)
	}

	defer respBody.Close()

	if statusCode != http.StatusOK {
		responseBytes, err := io.ReadAll(respBody)
		if err != nil {
			return fmt.Errorf("DownloadURLToPath: %w", err)
		}
		return fmt.Errorf("DownloadURLToPath: failed with (status: %v, response: %v)", statusCode, string(responseBytes))
	}

	_, err = io.Copy(out, respBody)
	if err != nil {
		return fmt.Errorf("DownloadURLToPath: failed to stream content from download URL to file: %w", err)
	}

	return nil
}

// Uses Get to download and PUT to upload
func DownloadAndUpload(
	ctx context.Context,
	downloadURL string,
	downloadHeaders map[string]string,
	uploadURL string,
	uploadHeaders map[string]string,
) error {
	if downloadURL == "" || uploadURL == "" {
		// nothing to do
		return nil
	}

	statusCode, downloadRespBody, err := DoHTTPStreamedCommunication(
		ctx,
		downloadURL,
		http.MethodGet,
		http.NoBody,
		downloadHeaders)
	if err != nil {
		return fmt.Errorf("DownloadAndUpload: %w", err)
	}

	defer downloadRespBody.Close()

	if statusCode != http.StatusOK {
		responseBytes, err := io.ReadAll(downloadRespBody)
		if err != nil {
			return fmt.Errorf("DownloadAndUpload: %w", err)
		}
		return fmt.Errorf("DownloadAndUpload: failed with (status: %v, response: %v)", statusCode, string(responseBytes))
	}

	statusCode, uploadRespBody, err := DoHTTPStreamedCommunication(
		ctx,
		uploadURL,
		http.MethodPut,
		downloadRespBody,
		uploadHeaders)
	if err != nil {
		return fmt.Errorf("DownloadAndUpload: %w", err)
	}

	defer uploadRespBody.Close()

	if statusCode != http.StatusOK {
		responseBytes, err := io.ReadAll(uploadRespBody)
		if err != nil {
			return fmt.Errorf("DownloadAndUpload: %w", err)
		}
		return fmt.Errorf("DownloadAndUpload: failed with (status: %v, response: %v)", statusCode, string(responseBytes))
	}

	fmt.Printf("DownloadAndUpload stats === \ndownloaded from - %v\nuploaded to - %v\n\n", downloadURL, uploadURL)

	return nil
}
