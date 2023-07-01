package database

import (
	"context"
	"fmt"
	"strconv"

	pb "web_hw1/pbGenerated"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgtype"
	"github.com/jackc/pgx/v5/pgxpool"
)

func ConnectToDatabase() (*pgxpool.Pool, error) {
	databaseUrl := "postgres://web_hw1:web_hw1@localhost:5432/postgres"
	conn, err := pgxpool.New(context.Background(), databaseUrl)
	if err != nil {
		return nil, fmt.Errorf("unable to connect to database: %w", err)
	}
	return conn, nil
}

func GetUsersFromDatabase(pool *pgxpool.Pool, userId int32) ([]*pb.User, error) {
	id := strconv.Itoa(int(userId))
	rows, err := pool.Query(context.Background(), "select * from users where id=$1", id)
	if err != nil {
		return nil, fmt.Errorf("get users failed: %w", err)
	}
	users, err := getUsersFromRows(rows)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	return users, nil
}

func GetFirst100UsersFromDatabase(pool *pgxpool.Pool) ([]*pb.User, error) {
	rows, err := pool.Query(context.Background(), "select * from users limit 100")
	if err != nil {
		return nil, fmt.Errorf("get users failed: %w", err)
	}
	users, err := getUsersFromRows(rows)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	return users, nil
}

func GetUsersWithInjectionFromDatabase(pool *pgxpool.Pool, userId string) ([]*pb.User, error) {
	query := "select * from users where id='" + userId + "'"
	rows, err := pool.Query(context.Background(), query)
	if err != nil {
		return nil, fmt.Errorf("get users failed: %w", err)
	}
	users, err := getUsersFromRows(rows)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	return users, nil
}

func getUsersFromRows(rows pgx.Rows) ([]*pb.User, error) {
	users := make([]*pb.User, 0)
	for rows.Next() {
		user := new(pb.User)
		var ts pgtype.Timestamp
		if err := rows.Scan(&user.Name, &user.Id, &user.Family, &user.Sex, &ts, &user.Age); err != nil {
			return nil, err
		}
		user.CreatedAt = ts.Time.String()
		users = append(users, user)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return users, nil
}
