package main

import (
	"database/sql"
	"fmt"
	"strconv"
	"time"

	"github.com/KristinaEtc/config"
	_ "github.com/KristinaEtc/slflog"
	"github.com/ventu-io/slf"
)

var log = slf.WithContext("main.go")

var (
	// These fields are populated by govvv
	BuildDate  string
	GitCommit  string
	GitBranch  string
	GitState   string
	GitSummary string
	Version    string
)

const (
	numOfComments int = 100
	numOfLikes    int = 50
)

type likeScheme struct {
	id                  int32
	comment_id          int
	user_id             int
	created_date        time.Time
	FK_LIKES_COMMENT_ID int
}

//Comments structure
type commentsScheme struct {
	id                int
	parent_id         int
	episode_id        int
	episode_timestamp int
	//type comment_type NOT NULL,
	user_id      int
	user_name    string
	body         string
	count_replay int
	count_likes  int
	created_date time.Time
}

type LikeConf struct {
	data       likeScheme
	numOfLikes int
}

type CommentConf struct {
	data       commentsScheme
	numOfLikes int
}

// ConfFile is a file with all program options
type ConfFile struct {
	Name     string
	Type     string
	User     string
	Password string
	NameDB   string
	Host     string
	Port     int
	Table    string
	Comments CommentConf
	Likes    LikeConf
}

var conf = ConfFile{
	Name:     "config",
	Type:     "postgres",
	User:     "guest",
	Password: "guest",
	NameDB:   "namedb",
	Host:     "127.0.0.1",
	Port:     5432,
	Table:    "comment",
}

func filDB(db *sql.DB) {

	var (
		sqlAddComment         = `INSERT INTO comment (user_name, comment, video_id, video_timestamp, calendar_timestamp)  VALUES ($1, $2, $3, $4, $5)`
		sqlAddLike            = `INSERT INTO comment (user_name, comment, video_id, video_timestamp, calendar_timestamp)  VALUES ($1, $2, $3, $4, $5)`
		sqlIncrementLikeCount = `INSERT INTO comment (user_name, comment, video_id, video_timestamp, calendar_timestamp)  VALUES ($1, $2, $3, $4, $5)`

		err         error
		i           int
		commentTemp = &commentsScheme{}
		likeTemp    = &likeScheme{}
	)

	for i = 0; i < numOfComments; i++ {
		commentTemp.body = "comment" + strconv.Itoa(i)
		commentTemp.count_likes = i  //random
		commentTemp.count_replay = i // random
		commentTemp.created_date = time.Now()
		commentTemp.episode_id = 13
		commentTemp.episode_timestamp = i // random
		commentTemp.id = i
		commentTemp.parent_id = 1 // random 1 or 0
		commentTemp.user_id = i   // random
		commentTemp.user_name = "anonim" + strconv.Itoa(i)

		_, err = db.Exec(sqlAddComment,
			commentTemp.body,
			commentTemp.count_likes,
			commentTemp.count_replay,
			commentTemp.created_date,
			commentTemp.episode_id,
			commentTemp.episode_timestamp,
			commentTemp.id,
			commentTemp.parent_id,
			commentTemp.user_id,
			commentTemp.user_name)
		if err != nil {
			log.Errorf("sqlAddComment: %s", err.Error())
			return
		}
	}

	for i = 0; i < numOfLikes; i++ {
		likeTemp.comment_id = i
		likeTemp.created_date = time.Now()
		likeTemp.FK_LIKES_COMMENT_ID = i
		likeTemp.id = 1
		likeTemp.user_id = 1

		_, err = db.Exec(sqlAddLike,
			likeTemp.comment_id,
			likeTemp.created_date,
			likeTemp.FK_LIKES_COMMENT_ID,
			likeTemp.id,
			likeTemp.user_id)
		if err != nil {
			log.Errorf("sqlAddLike: %s", err.Error())
			return
		}

		_, err = db.Exec(sqlIncrementLikeCount,
			likeTemp.comment_id,
			likeTemp.created_date,
			likeTemp.FK_LIKES_COMMENT_ID,
			likeTemp.id,
			likeTemp.user_id)
		if err != nil {
			log.Errorf("sqlIncrementLikeCount: %s", err.Error())
			return
		}
	}
}

func initDB() (db *sql.DB, err error) {
	sqlOpenStr := fmt.Sprintf("host=%s port=%d user=%s "+"password=%s dbname=%s sslmode=disable",
		conf.Host, conf.Port, conf.User, conf.Password, conf.NameDB)
	log.Info(sqlOpenStr)

	db, err = sql.Open("postgres", sqlOpenStr)
	if err != nil {
		log.Errorf("open database: %s", err.Error())
		return nil, fmt.Errorf("cannot opet database: %s", err.Error())
	}

	if err = db.Ping(); err != nil {
		log.Errorf("ping database: %s", err.Error())
		return nil, err
	}
	return
}

func main() {
	config.ReadGlobalConfig(&conf, "main")
	var (
		err error
		db  *sql.DB
	)
	db, err = initDB()
	if err != nil {
		return
	}
	filDB(db)
}
