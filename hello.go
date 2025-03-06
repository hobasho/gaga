package main

//помарочки были\\чисто ради коммита...2
import (
	"encoding/json"
	"fmt"
	"net/http"
	"sort"
)

// Структуры для парсинга JSON ответа
type PackageInfo struct {
	Name      string `json:"name"`
	Epoch     int    `json:"epoch"`
	Version   string `json:"version"`
	Release   string `json:"release"`
	Arch      string `json:"arch"`
	Disttag   string `json:"disttag"`
	Buildtime int    `json:"buildtime"`
	Source    string `json:"source"`

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
		fmt.Printf("HTTP error : %s\n", resp.Status)
		return
	}

	var result APIResponse
	decoder := json.NewDecoder(resp.Body)
	if err := decoder.Decode(&result); err != nil {
		fmt.Printf("JSON decode error: %v\n", err)
		return
	}

	// Собираем уникальные архитектуры
	archSet := make(map[string]struct{})
	for _, pkg := range result.Packages {
		if pkg.Arch != "" { // Игнорируем пустые значения
			archSet[pkg.Arch] = struct{}{}
		}
	}

	// Преобразуем в срез и сортируем
	architectures := make([]string, 0, len(archSet))
	for arch := range archSet {
		architectures = append(architectures, arch)
	}
	sort.Strings(architectures)

	// Выводим результат
	fmt.Printf("Found %d unique architectures in branch '%s':\n",
		len(architectures), result.Branch)
	for _, arch := range architectures {
		fmt.Println(arch)
	}

}
