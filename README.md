# abci-demo

```bash
cd $GOPATH/github.com/
mkdir tendermint
cd tendermint
git clone https://github.com/tendermint/tendermint.git
cd tendermint
mkdir apps
cd apps
git clone https://github.com/chinajsstar/abci-demo.git

cd ..
cd proxy

#修改client.go文件，添加刚刚apps里面的abci-demo包
#在DefaultClientCreator里面加上一个case选项如下：
#case "abci-demo":
#    return NewLocalClientCreator(abci-demo.NewPersistentABCIDemoApplication(dbDir))

#vim client.go

```
