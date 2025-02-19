package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// Структуры для парсинга JSON ответа
type PackageInfo struct {
	Name    string `json:"name"`
	Version string `json:"version"`
	Arch    string `json:"arch"`
	// Добавьте другие нужные поля по документации API
}

type APIResponse struct {
	Branch   string        `json:"branch"`
	Packages []PackageInfo `json:"packages"`
	// Дополнительные поля ответа
}

func main() {
	url := "https://rdb.altlinux.org/api/export/branch_binary_packages/sisyphus"

	resp, err := http.Get(url)
	if err != nil {
		fmt.Printf("Request error: %v\n", err)
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		fmt.Printf("HTTP error: %s\n", resp.Status)
		return
	}

	var result APIResponse
	decoder := json.NewDecoder(resp.Body)
	if err := decoder.Decode(&result); err != nil {
		fmt.Printf("JSON decode error: %v\n", err)
		return
	}

	// Выводим имена пакетов
	fmt.Printf("Found %d packages in branch '%s':\n", len(result.Packages), result.Branch)
	for _, pkg := range result.Packages {
		fmt.Println(pkg.Name)
	}

}
