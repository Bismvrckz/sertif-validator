package databases

import "tkbai-be/config"

func ConnectTkbaiDatabase() (err error) {
	cmsDB, err := config.TkbaiDbConnection()
	if err != nil {
		return err
	}

	err = cmsDB.Ping()
	if err != nil {
		return err
	}

	return nil
}
