5.优雅退出
方式1：直接指定配置文件路径（相对路径或者绝对路径）
相对路径：相对执行的可执行文件的相对路径
viper.SetConfigFile("./conf/config.yaml")
绝对路径：系统中实际的文件路径
viper.SetConfigFile("/Users/liwenzhou/Desktop/bluebell/conf/config.yaml")

方式2：指定配置文件名和配置文件的位置，viper自行查找可用的配置文件
配置文件名不需要带后缀
配置文件位置可配置多个
viper.SetConfigName("config") // 指定配置文件名（不带后缀）
viper.AddConfigPath(".") // 指定查找配置文件的路径（这里使用相对路径）
viper.AddConfigPath("./conf")      // 指定查找配置文件的路径（这里使用相对路径）

基本上是配合远程配置中心使用的，告诉viper当前的数据使用什么格式去解析
viper.SetConfigType("json")   