package config

import db "backend/database"

type config struct {
	MajorRequirements map[string][][]db.Class
}
