# go-yudan
---------------------
Get the bullet screen of the site(douyu.com) <br>

[👉 Java version](https://github.com/AlexJialene/yudan) <br>
[👉 Document description](https://github.com/AlexJialene/yudan/blob/master/README.md)

## Example

```$xslt
go get github.com/AlexJialene/go-yudan

```

```$xslt

// roomId , groupId , callback func
bullet.Start("24422", "-9999", func(msg map[string]string) {

		fmt.Println(msg["type"])
		//TODO

})
```


[See the documentation for details / View details of callback collections ](http://dev-bbs.douyutv.com/forum.php?mod=attachment&aid=MjkxfDQ5ZGQ5NmUwfDE1NTkyNzQ2MTh8MHwzOTk%3D)



