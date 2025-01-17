package db

import (
	"context"
	"errors"
	"simple-tiktok/pkg/consts"

	"gorm.io/gorm"
)

type Follow struct {
	gorm.Model
	UserID       uint   `json:"user_id"`
	Username     string `json:"username"`
	FollowerID   uint   `json:"follower_id"`
	FollowerName string `json:"follower_name"`
}

func (f *Follow) TableName() string {
	return consts.FollowTableName
}

// 查询uid的粉丝
func QueryFollower(ctx context.Context, uid uint) ([]*Follow, error) {
	followers := make([]*Follow, 0)
	err := DB.WithContext(ctx).Where("user_id = ?", uid).Find(&followers).Error
	return followers, err
}

// 查询uid的关注
func QueryFollow(ctx context.Context, uid uint) ([]*Follow, error) {
	follows := make([]*Follow, 0)
	err := DB.WithContext(ctx).Where("follower_id = ?", uid).Find(&follows).Error
	return follows, err
}

// uid1关注uid2
func FollowUser(ctx context.Context, uid1 uint, uid1Name string, uid2 uint, uid2Name string) error {
	// 先查软删除的记录
	result := DB.WithContext(ctx).Unscoped().Where("user_id = ? and follower_id = ?", uid2, uid1).Take(&Follow{})
	// 没有软删除的记录就新加一条记录
	if err := result.Error; errors.Is(err, gorm.ErrRecordNotFound) {
		return DB.WithContext(ctx).Create(&Follow{
			UserID:       uid2,
			Username:     uid2Name,
			FollowerID:   uid1,
			FollowerName: uid1Name,
		}).Error
	} else if err != nil {
		return err
	}
	// 如果有，就更新deleted_at值
	return result.Update("deleted_at", nil).Error
}

// uid1取关uid2
func UnFollowUser(ctx context.Context, uid1, uid2 uint) error {
	return DB.WithContext(ctx).Where("user_id = ? and follower_id = ?", uid2, uid1).Delete(&Follow{}).Error
}

// 查询uids列表中被uid关注
func MGetFollow(ctx context.Context, uid uint, uids []uint) ([]uint, error) {
	res := make([]uint, 0, len(uids))
	err := DB.WithContext(ctx).Model(&Follow{}).Select("user_id").Where("user_id in ? and follower_id = ?", uids, uid).Find(&res).Error
	return res, err
}
