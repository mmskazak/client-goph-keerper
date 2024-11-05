package commands

import (
	"client-goph-keerper/internal/storage"
	"fmt"
	"io"
	"log"
	"net/http"
	"path"

	"github.com/spf13/cobra"
)

// SetDeleteFileCmd создает команду для удаления файла по ID.
func SetDeleteFileCmd(s *storage.Storage) (*cobra.Command, error) {
	deleteFileCmd := &cobra.Command{
		Use:   "delete",
		Short: "Delete a file by ID",
		RunE: func(cmd *cobra.Command, args []string) error {
			fileID, err := cmd.Flags().GetString("file_id")
			if err != nil {
				return fmt.Errorf("file id is required: %w", err)
			}

			reqURL := path.Join(s.ServerURL, "file", "delete", fileID)
			req, err := http.NewRequest(http.MethodGet, reqURL, http.NoBody)
			if err != nil {
				return fmt.Errorf("ошибка создания запроса: %w", err)
			}

			req.Header.Set("Content-Type", "application/json")
			req.Header.Set("Authorization", s.Token)

			client := &http.Client{}
			resp, err := client.Do(req)
			if err != nil {
				return fmt.Errorf("ошибка отправки запроса: %v", err)
			}
			defer func(Body io.ReadCloser) {
				err := Body.Close()
				if err != nil {
					log.Printf("error closing response body: %v", err)
				}
			}(resp.Body)

			fmt.Printf("Response: %v\n", resp.Status)
			return nil
		},
	}

	deleteFileCmd.Flags().String("file_id", "", "File ID to delete")
	err := deleteFileCmd.MarkFlagRequired("file_id")
	if err != nil {
		return nil, fmt.Errorf("ошибка установки обязательного флага 'file_id': %v", err)
	}

	return deleteFileCmd, nil
}
