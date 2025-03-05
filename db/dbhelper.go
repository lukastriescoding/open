package db

import (
	"database/sql"
	"errors"
	"fmt"
	"os"
	"path/filepath"

	"github.com/lukastriescoding/open/models"
	"github.com/mattn/go-sqlite3"
)

var (
	conn   *sql.DB
	dbdir  string
	dbname string
	dbpath string
)

func init() {
	var homedir, err = os.UserHomeDir()
	if err != nil {
		panic(err)
	}
	dbdir = homedir + "/.open"
	dbname = "open.db"
	dbpath = dbdir + "/" + dbname
}

func InitCon() error {
	var err error
	_, err = os.Stat(dbpath)
	if os.IsNotExist(err) {
		return createDB(true)
	}
	conn, err = sql.Open("sqlite3", dbpath)
	if err != nil {
		return err
	}
	err = verifyDBSchema()
	if err != nil {
		return err
	}
	return nil
}

func InsertDir(name, path string) error {
	_, err := conn.Exec("INSERT INTO directories (name, path) VALUES (?, ?)", name, path)
	if sqliteErr, ok := err.(sqlite3.Error); ok && sqliteErr.Code == sqlite3.ErrConstraint {
		return errors.New("Directory with name \"" + name + "\" already exists")
	}
	return err
}

func GetDir(name string) (models.Directory, error) {
	var dir models.Directory
	row := conn.QueryRow("SELECT name, path, main_app FROM directories WHERE name = ?", name)
	if row.Err() != nil {
		return models.Directory{}, row.Err()
	}
	err := row.Scan(&dir.Name, &dir.Path, &dir.MainApp)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			// Handle the case where no rows were found
			return models.Directory{}, errors.New("directory not found") // or a custom error if needed
		}
		return models.Directory{}, err
	}
	return dir, nil
}

func GetAllDirs() ([]models.Directory, error) {
	rows, err := conn.Query("SELECT name, path, main_app FROM directories")
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var dirs []models.Directory
	for rows.Next() {
		var dir models.Directory
		err = rows.Scan(&dir.Name, &dir.Path, &dir.MainApp)
		if err != nil {
			return nil, err
		}
		dirs = append(dirs, dir)
	}
	return dirs, nil
}

func UpdateDirMainApp(dirName, appName string) error {
	result, err := conn.Exec("UPDATE directories SET main_app = ? WHERE name = ?", appName, dirName)
	if err != nil {
		return err
	}
	rows, err := result.RowsAffected()
	if err != nil {
		return err
	}
	if rows == 0 {
		return errors.New("directory with name \"" + dirName + "\" not found")
	}
	return err
}

func RemoveDir(name string) error {
	result, err := conn.Exec("DELETE FROM directories WHERE name = ?", name)
	if err != nil {
		return err
	}
	rows, err := result.RowsAffected()
	if err != nil {
		return err
	}
	if rows == 0 {
		return errors.New("directory with name \"" + name + "\" not found")
	}
	return err
}

func InsertApp(name, path string) error {
	if name == "" {
		name = filepath.Base(path)
	}
	_, err := conn.Exec("INSERT INTO applications (name, path) VALUES (?, ?)", name, path)
	if sqliteErr, ok := err.(sqlite3.Error); ok && sqliteErr.Code == sqlite3.ErrConstraint {
		return errors.New("App with name \"" + name + "\" already exists")
	}
	return err
}

func ExistsApp(name string) (bool, error) {
	var exists bool
	err := conn.QueryRow("SELECT EXISTS(SELECT 1 FROM applications WHERE name = ?)", name).Scan(&exists)
	if err != nil {
		return false, err
	}
	return exists, nil
}

func GetApp(name string) (models.Application, error) {
	var app models.Application
	row := conn.QueryRow("SELECT name, path FROM applications WHERE name = ?", name)
	if row.Err() != nil {
		return models.Application{}, row.Err()
	}
	err := row.Scan(&app.Name, &app.Path)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			// Handle the case where no rows were found
			return models.Application{}, errors.New("application not found") // or a custom error if needed
		}
		return models.Application{}, err
	}
	return app, nil
}

func GetAllApps() ([]models.Application, error) {
	rows, err := conn.Query("SELECT name, path FROM applications")
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var apps []models.Application
	for rows.Next() {
		var app models.Application
		err = rows.Scan(&app.Name, &app.Path)
		if err != nil {
			return nil, err
		}
		apps = append(apps, app)
	}
	return apps, nil
}

func RemoveApp(name string) error {
	result, err := conn.Exec("DELETE FROM applications WHERE name = ?", name)
	if err != nil {
		return err
	}
	rows, err := result.RowsAffected()
	if err != nil {
		return err
	}
	if rows == 0 {
		return errors.New("application with name \"" + name + "\" not found")
	}
	return err
}

func SetMainApp(name string) error {
	_, err := conn.Exec("DELETE FROM main_app")
	if err != nil {
		return err
	}
	_, err = conn.Exec("INSERT INTO main_app (name) VALUES (?)", name)
	if err != nil {
		return err
	}
	return nil
}

func GetMainApp() (models.Application, error) {
	// TODO
	var name string
	var app models.Application
	err := conn.QueryRow("SELECT name FROM main_app").Scan(&name)
	if err != nil {
		return models.Application{}, err
	}

	return app, nil
}

func Close() {
	conn.Close()
}

func verifyDBSchema() error {
	var dirSchemaWant = "CREATE TABLE directories (name TEXT PRIMARY KEY, path TEXT, main_app TEXT, FOREIGN KEY (main_app) REFERENCES applications(name) ON DELETE CASCADE)"
	var appSchemaWant = "CREATE TABLE applications (name TEXT PRIMARY KEY, path TEXT)"
	var mainappSchemaWant = "CREATE TABLE main_app (name TEXT PRIMARY KEY, FOREIGN KEY (name) REFERENCES applications(name) ON DELETE CASCADE)"
	var dirSchemaGot string
	query := `SELECT sql FROM sqlite_master WHERE type='table' AND name='directories'`
	err := conn.QueryRow(query).Scan(&dirSchemaGot)
	if err != nil {
		return structureDB()
	}
	if dirSchemaGot != dirSchemaWant {
		return structureDB()
	}
	var appSchemaGot string
	query = `SELECT sql FROM sqlite_master WHERE type='table' AND name='applications'`
	err = conn.QueryRow(query).Scan(&appSchemaGot)
	if err != nil {
		return structureDB()
	}
	if appSchemaGot != appSchemaWant {
		return structureDB()
	}
	var mainappSchemaGot string
	query = `SELECT sql FROM sqlite_master WHERE type='table' AND name='main_app'`
	err = conn.QueryRow(query).Scan(&mainappSchemaGot)
	if err != nil {
		return structureDB()
	}
	if mainappSchemaGot != mainappSchemaWant {
		return structureDB()
	}
	return nil
}

func structureDB() error {
	fmt.Println("Database seems to exist, but with a wrong schema. Want to fix it automatically?\n(The DB will be deleted and recreated) (y/n):")
	var answer string
	fmt.Scanln(&answer)
	if answer != "y" {
		return errors.New("cancelling: db not updated")
	}
	err := os.Remove(dbpath)
	if err != nil {
		return err
	}
	err = createDB(false)
	if err != nil {
		return err
	}
	return nil
}

func createDB(ask bool) error {
	if ask {
		fmt.Println("Database not found under " + dbpath + ". Want to create a new one? (y/n):")
		var answer string
		fmt.Scanln(&answer)
		if answer != "y" {
			return errors.New("cancelling: db not created")
		}
	}

	err := os.MkdirAll(dbdir, 0755)
	if err != nil {
		return fmt.Errorf("failed to create db-directory: %w", err)
	}

	file, err := os.Create(dbpath)
	if err != nil {
		return err
	}
	file.Close()
	conn, err = sql.Open("sqlite3", dbpath)
	if err != nil {
		return err
	}
	err = createDBStructur()
	if err != nil {
		return err
	}
	fmt.Println("Database created under " + dbpath)
	return nil
}

func createDBStructur() error {
	_, err := conn.Exec("CREATE TABLE IF NOT EXISTS directories (name TEXT PRIMARY KEY, path TEXT, main_app TEXT, FOREIGN KEY (main_app) REFERENCES applications(name) ON DELETE CASCADE)")
	if err != nil {
		return err
	}
	_, err = conn.Exec("CREATE TABLE IF NOT EXISTS applications (name TEXT PRIMARY KEY, path TEXT)")
	if err != nil {
		return err
	}

	_, err = conn.Exec("CREATE TABLE IF NOT EXISTS main_app (name TEXT PRIMARY KEY, FOREIGN KEY (name) REFERENCES applications(name) ON DELETE CASCADE)")
	if err != nil {
		return err
	}
	return nil
}
