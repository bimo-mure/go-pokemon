package models

import (
	"context"
	"database/sql"
	"time"

	"github.com/dev-bimomure/go-pokemon/app/structs"
)

type DBModel struct {
	DB *sql.DB
}

type Models struct {
	DB DBModel
}

func NewModels(db *sql.DB) Models {
	return Models{
		DB: DBModel{DB: db},
	}
}

func (m *DBModel) GetPokemon() ([]structs.MyPokemon, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	var listPokemon []structs.MyPokemon

	rows, err := m.DB.QueryContext(ctx, "SELECT id, pokemon_name, nick_name, image FROM my_pokemon")
	for rows.Next() {
		var myPokemon structs.MyPokemon
		err = rows.Scan(&myPokemon.Id, &myPokemon.PokemonName, &myPokemon.NickName, &myPokemon.Image)
		if err != nil {
			return listPokemon, err
		}
		listPokemon = append(listPokemon, myPokemon)
	}

	if err := rows.Err(); err != nil {
		return listPokemon, err
	}
	return listPokemon, err
}

func (m *DBModel) InsertPokemon(pkm structs.MyPokemon) (int, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	insert := `INSERT INTO my_pokemon (pokemon_name, nick_name, image) values (?,?,?)`

	result, err := m.DB.ExecContext(ctx, insert,
		pkm.PokemonName,
		pkm.NickName,
		pkm.Image,
	)
	if err != nil {
		return 0, err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}

	return int(id), nil
}

func (m *DBModel) DeletePokemon(pokemonId int) (int, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	delete := `DELETE FROM my_pokemon WHERE id = ?`

	result, err := m.DB.ExecContext(ctx, delete, pokemonId)
	if err != nil {
		return 0, err
	}

	rowDeleted, err := result.RowsAffected()
	if err != nil {
		return 0, err
	}

	return int(rowDeleted), nil
}

func (m *DBModel) UpdatePokemon(pokemonId int, newNickname string, updateCount int) (int, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	updatePoke := `UPDATE my_pokemon SET nick_name = ?, update_count = ? WHERE id = ?`

	result, err := m.DB.ExecContext(ctx, updatePoke, newNickname, updateCount, pokemonId)
	if err != nil {
		return 0, err
	}

	rowUpdated, err := result.RowsAffected()
	if err != nil {
		return 0, err
	}

	return int(rowUpdated), nil
}

func (m *DBModel) GetPokemonLastCountById(pokemonId int) (int, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	var count int
	stmt := `SELECT update_count FROM my_pokemon WHERE id = ?`
	row := m.DB.QueryRowContext(ctx, stmt, pokemonId)
	err := row.Scan(&count)
	if err != nil {
		return count, err
	}

	return count, nil
}
