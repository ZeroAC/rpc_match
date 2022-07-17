namespace go match_service //声明转换为go语言

struct User {   //定义结构体 User
    1: i32 id, 
    2: string name,
    3: i32 score
}
//生成服务代码
//kitex -module github.com/ZeroAC/rpc_match -service match_service  ./thrift/match_service.thrift
service Match { //定义service服务
    /** 
     * user: 添加的用户信息
     * info: 附加信息
     * 在匹配池中添加一个名用户
     */  

    i32 add_user(1: User user, 2: string info),

    /** 
     * user: 删除的用户信息
     * info: 附加信息
     * 从匹配池中删除一名用户
     */  

    i32 remove_user(1: User user, 2: string info),
}