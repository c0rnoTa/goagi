package mysql

import (
	"context"
	"errors"
	"fmt"
)

func (s *Storage) BlacklistCheck(ctx context.Context, caller string) (int, error) {
	var blacklist int

	query := `SELECT 1 FROM black_list WHERE 
                             num LIKE ? OR 
                             num LIKE CASE WHEN LENGTH(?)=11 AND (substr(?,1,1) IN (7,8)) THEN CONCAT('_',SUBSTR(?,-10)) ELSE FALSE END OR 
                             num LIKE CASE WHEN LENGTH(?)=11 AND (substr(?,1,1) IN (7,8)) THEN SUBSTR(?,-10) ELSE FALSE END`
	if err := s.checkConnection(ctx); err != nil {
		return 0, errors.New("could not connect to database")
	}

	if err := s.DB.QueryRow(query, caller, caller, caller, caller, caller, caller, caller).Scan(&blacklist); err != nil {
		return 0, fmt.Errorf("unable to query row in database: %w", err)
	}

	return blacklist, nil
}
