package sqlite

import (
	"database/sql"
	"time"

	"github.com/coline-carle/zhevra/storage/model"
	"github.com/pkg/errors"
)

// CreateCurseReleaseDirectories create as many folder as we have
// it does not return an error if the folder already exist
// return a list of id for the coressponding folders
func (s *Storage) CreateCurseReleaseDirectories(
	tx *sql.Tx, release model.CurseRelease) error {
	for _, directory := range release.Directories {
		_, err := tx.Exec(`
				INSERT INTO curse_release_directory
					(release_id, directory)
				VALUES(?, ?)`,
			release.ID, directory)
		if err != nil {
			return err
		}
	}
	return nil
}

// CreateCurseReleaseGameVersions create as many folder as we have
// it does not return an error if the folder already exist
// return a list of id for the coressponding folders
func (s *Storage) CreateCurseReleaseGameVersions(
	tx *sql.Tx, release model.CurseRelease) error {
	for _, gameVersion := range release.GameVersions {
		_, err := tx.Exec(`
				INSERT INTO curse_release_game_version
					(release_id, game_version)
				VALUES(?, ?)`,
			release.ID, gameVersion)
		if err != nil {
			return err
		}
	}
	return nil
}

// CreateCurseRelease save new addon in the database
func (s *Storage) CreateCurseRelease(
	tx *sql.Tx, release model.CurseRelease) error {
	_, err := tx.Exec(
		`INSERT INTO curse_release (
			id,
			filename,
			created_at,
			url,
			addon_id,
			is_alternate
		) VALUES(?, ?, ?, ?, ?, ?)
		`,
		release.ID,
		release.Filename,
		release.CreatedAt.Unix(),
		release.URL,
		release.AddonID,
		release.IsAlternate,
	)
	if err != nil {
		return errors.Wrap(err, "INSERT curse_release failed")
	}
	err = s.CreateCurseReleaseGameVersions(tx, release)
	if err != nil {
		return errors.Wrap(err, "failed to CreateCurseRelease (GameVersion)")
	}
	err = s.CreateCurseReleaseDirectories(tx, release)
	if err != nil {
		return errors.Wrap(err, "failed to CreateCurseRelease (Directories)")
	}
	return nil
}

// FindCurseReleaseDirectoriesByReleaseID return all directories for a
// given release
func (s *Storage) FindCurseReleaseDirectoriesByReleaseID(
	tx *sql.Tx, id int64) ([]string, error) {
	directories := []string{}
	rows, err := tx.Query(`
		SELECT
			directory
		FROM
			curse_release_directory
		WHERE
			release_id = $1
		`, id)
	if err != nil {
		return directories, errors.Wrap(err, "FindCurseReleaseDirecotries failed")
	}
	defer rows.Close()
	return rowsToStringSlice(rows)
}

// FindCurseReleaseGameVersionsByReleaseID return all gameversions for a
// given release
func (s *Storage) FindCurseReleaseGameVersionsByReleaseID(
	tx *sql.Tx, id int64) ([]int, error) {
	gameVersions := []int{}
	rows, err := tx.Query(`
		SELECT
			game_version
		FROM
			curse_release_game_version
		WHERE
			release_id = $1
		`, id)
	if err != nil {
		return gameVersions, errors.Wrap(err, "FindCurseReleaseGameVersions failed")
	}
	defer rows.Close()
	return rowsToIntSlice(rows)
}

// FindCurseReleasesByAddonID return all release for a given addon ID
func (s *Storage) FindCurseReleasesByAddonID(
	tx *sql.Tx, id int64) ([]model.CurseRelease, error) {
	releases := []model.CurseRelease{}
	rows, err := tx.Query(`
		SELECT
			id,
			filename,
			created_at,
			url,
			addon_id,
			is_alternate
		FROM
			curse_release
		WHERE
			addon_id = $1
		`, id)
	if err != nil {
		return releases, errors.Wrap(err, "FindCurseReleasesByAddonID")
	}
	defer rows.Close()
	return s.rowsToReleases(tx, rows)
}

func rowsToIntSlice(rows *sql.Rows) ([]int, error) {
	ints := []int{}
	var intValue int
	for rows.Next() {
		err := rows.Scan(&intValue)
		if err != nil {
			return ints, errors.Wrap(err, "rowsToIntSlice")
		}
		ints = append(ints, intValue)
	}
	return ints, nil
}

func rowsToStringSlice(rows *sql.Rows) ([]string, error) {
	directories := []string{}
	var directory string
	for rows.Next() {
		err := rows.Scan(&directory)
		if err != nil {
			return directories, errors.Wrap(err, "rowsToStringSlice")
		}
		directories = append(directories, directory)
	}
	return directories, nil
}

func (s *Storage) rowsToReleases(
	tx *sql.Tx, rows *sql.Rows) ([]model.CurseRelease, error) {
	releases := []model.CurseRelease{}
	var date int64
	for rows.Next() {
		release := model.CurseRelease{}
		err := rows.Scan(
			&release.ID,
			&release.Filename,
			&date,
			&release.URL,
			&release.AddonID,
			&release.IsAlternate,
		)
		if err != nil {
			return releases, errors.Wrap(err, "rowsToReleases")
		}
		release.CreatedAt = time.Unix(date, 0).UTC()
		release.GameVersions, err = s.FindCurseReleaseGameVersionsByReleaseID(tx, release.ID)
		if err != nil {
			return releases, errors.Wrap(err, "rowsToReleases")
		}
		release.Directories, err = s.FindCurseReleaseDirectoriesByReleaseID(tx, release.ID)
		releases = append(releases, release)
	}
	return releases, nil
}

// GetCurseRelease fetch and addon by ID
func (s *Storage) GetCurseRelease(
	tx *sql.Tx, id int64) (model.CurseRelease, error) {
	release := model.CurseRelease{}
	var date int64
	err := tx.QueryRow(`
		SELECT
			id,
			filename,
			created_at,
			url,
			addon_id,
			is_alternate
		FROM
			curse_release
		WHERE
			id = $1
		`, id).Scan(
		&release.ID,
		&release.Filename,
		&date,
		&release.URL,
		&release.AddonID,
		&release.IsAlternate,
	)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return release, errors.Wrapf(err, "id %d", id)
		}
		return release, errors.Wrap(err, "GetCurseRelease failed")
	}
	release.Directories, err = s.FindCurseReleaseDirectoriesByReleaseID(tx, release.ID)
	if err != nil {
		return release, errors.Wrap(err, "GetCurseRelease failed")
	}
	release.GameVersions, err = s.FindCurseReleaseGameVersionsByReleaseID(tx, release.ID)
	release.CreatedAt = time.Unix(date, 0).UTC()
	return release, err
}
