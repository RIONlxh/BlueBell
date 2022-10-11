package mysql

import (
	"BlueBell/models"
	"fmt"
)

func CreateOnePost(p *models.Post) (err error) {
	sqlStr := `insert into post(
    post_id, user_id, community_id, title, content)
	values (?, ?, ?, ?, ?)
	`
	fmt.Println(p)
	_, err = db.Exec(sqlStr, p.PostID, p.UserID, p.CommunityID, p.Title, p.Content)
	if err != nil {
		return err
	}
	return
}

func DeleteOnePost(post_id int64) (err error) {
	sql := "delete from post where id = ?"
	_, err = db.Exec(sql, post_id)
	if err != nil {
		return err
	}
	return nil
}
