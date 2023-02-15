// Code generated by hertz generator. DO NOT EDIT.

package Tiktok

import (
	"github.com/cloudwego/hertz/pkg/app/server"
	tiktok "simple-tiktok/biz/handler/tiktok"
)

/*
 This file will register all the routes of the services in the master idl.
 And it will update automatically when you use the "update" command for the idl.
 So don't modify the contents of the file, or your code will be deleted when it is updated.
*/

// Register register routes based on the IDL 'api.${HTTP Method}' annotation.
func Register(r *server.Hertz) {

	root := r.Group("/", rootMw()...)
	{
		_douyin := root.Group("/douyin", _douyinMw()...)
		{
			_comment := _douyin.Group("/comment", _commentMw()...)
			{
				_action := _comment.Group("/action", _actionMw()...)
				_action.POST("/", append(_uploadcommentMw(), tiktok.UploadComment)...)
			}
		}
		{
			_favorite := _douyin.Group("/favorite", _favoriteMw()...)
			{
				_action0 := _favorite.Group("/action", _action0Mw()...)
				_action0.POST("/", append(_favourite_ctionMw(), tiktok.FavouriteAction)...)
			}
			{
				_list := _favorite.Group("/list", _listMw()...)
				_list.POST("/", append(_getfavouritelistMw(), tiktok.GetFavouriteList)...)
			}
		}
		{
			_feed := _douyin.Group("/feed", _feedMw()...)
			_feed.GET("/", append(_feed0Mw(), tiktok.Feed)...)
		}
		{
			_publish := _douyin.Group("/publish", _publishMw()...)
			{
				_action1 := _publish.Group("/action", _action1Mw()...)
				_action1.POST("/", append(_uploadvideoMw(), tiktok.UploadVideo)...)
			}
		}
		{
			_relation := _douyin.Group("/relation", _relationMw()...)
			{
				_action2 := _relation.Group("/action", _action2Mw()...)
				_action2.POST("/", append(_followuserMw(), tiktok.FollowUser)...)
			}
			{
				_follow := _relation.Group("/follow", _followMw()...)
				{
					_list0 := _follow.Group("/list", _list0Mw()...)
					_list0.GET("/", append(_getfollowMw(), tiktok.GetFollow)...)
				}
			}
			{
				_follower := _relation.Group("/follower", _followerMw()...)
				{
					_list1 := _follower.Group("/list", _list1Mw()...)
					_list1.GET("/", append(_getfollowerMw(), tiktok.GetFollower)...)
				}
			}
		}
		{
			_user := _douyin.Group("/user", _userMw()...)
			_user.GET("/", append(_getuserMw(), tiktok.GetUser)...)
			{
				_login := _user.Group("/login", _loginMw()...)
				_login.POST("/", append(_checkuserMw(), tiktok.CheckUser)...)
			}
			{
				_register := _user.Group("/register", _registerMw()...)
				_register.POST("/", append(_createuserMw(), tiktok.CreateUser)...)
			}
		}
	}
}
