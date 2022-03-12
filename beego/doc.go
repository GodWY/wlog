// 主要的参数如下说明(除 separate 外,均与 file 相同)：

// filename 保存的文件名
// maxlines 每个文件保存的最大行数，默认值 1000000
// maxsize 每个文件保存的最大尺寸，默认值是 1 << 28, //256 MB
// daily 是否按照每天 logrotate，默认是 true
// maxdays 文件最多保存多少天，默认保存 7 天
// rotate 是否开启 logrotate，默认是 true
// level 日志保存的时候的级别，默认是 Trace 级别
// perm 日志文件权限
// separate 需要单独写入文件的日志级别,设置后命名类似 test.error.log

package beego
