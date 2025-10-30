package db

import (
	"database/sql"
	"log"
	"os"

	_ "github.com/lib/pq"
)

func ConnectDB() (*sql.DB, error) {
	// Pega a URL do banco do Render
	dbURL := os.Getenv("DATABASE_URL")
	if dbURL == "" {
		log.Println("❌ ERRO: DATABASE_URL não definida")
		return nil, sql.ErrConnDone
	}

	// Conecta no banco usando a URL completa
	db, err := sql.Open("postgres", dbURL)
	if err != nil {
		log.Println("❌ Erro ao abrir conexão com BD:", err)
		return nil, err
	}

	// Testa conexão
	err = db.Ping()
	if err != nil {
		log.Println("❌ Não conectou ao BD:", err)
		return nil, err
	}

	log.Println("✅ Conectado ao banco com sucesso!")
	return db, nil
}
