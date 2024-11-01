package commands

import (
	"bytes"
	"client-goph-keerper/internal/storage"
	"encoding/json"
	"fmt"
	"github.com/spf13/cobra"
	"net/http"
)

// SetDeleteFileCmd создает команду для удаления файла по ID
func SetDeleteFileCmd(s *storage.Storage) (*cobra.Command, error) {
	deleteFileCmd := &cobra.Command{
		Use:   "delete",
		Short: "Delete a file by ID",
		RunE: func(cmd *cobra.Command, args []string) error {
			fileID, _ := cmd.Flags().GetString("file_id")
			userID, _ := cmd.Flags().GetInt("user_id")

			data := map[string]interface{}{
				"file_id": fileID,
				"user_id": userID,
			}

			body, err := json.Marshal(data)
			if err != nil {
				return fmt.Errorf("ошибка кодирования JSON: %v", err)
			}

			req, err := http.NewRequest("POST", fmt.Sprintf("%s/file/delete", s.ServerURL), bytes.NewBuffer(body))
			if err != nil {
				return fmt.Errorf("ошибка создания запроса: %v", err)
			}

			req.Header.Set("Content-Type", "application/json")
			req.Header.Set("Authorization", s.Token)

			client := &http.Client{}
			resp, err := client.Do(req)
			if err != nil {
				return fmt.Errorf("ошибка отправки запроса: %v", err)
			}
			defer resp.Body.Close()

			fmt.Printf("Response: %v\n", resp.Status)
			return nil
		},
	}

	deleteFileCmd.Flags().String("file_id", "", "File ID to delete")
	deleteFileCmd.Flags().Int("user_id", 0, "User ID")
	err := deleteFileCmd.MarkFlagRequired("file_id")
	if err != nil {
		return nil, fmt.Errorf("ошибка установки обязательного флага 'file_id': %v", err)
	}
	err = deleteFileCmd.MarkFlagRequired("user_id")
	if err != nil {
		return nil, fmt.Errorf("ошибка установки обязательного флага 'user_id': %v", err)
	}
	err = deleteFileCmd.MarkFlagRequired("token")
	if err != nil {
		return nil, fmt.Errorf("ошибка установки обязательного флага 'token': %v", err)
	}

	return deleteFileCmd, nil
}
