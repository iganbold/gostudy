## Rss reader (Written in Go)
### Description
```
The app reads a json based rss url data from a given file then unmarshall it. After that, it sends the http GET request basend on the urls to fetch the news then it saves the http response to custom struct type. Finally, it prints the web site name and the news title.

* USED DATASTRUCTURES: slice, struct,
* USED IO LIBS: os, http
* USED ENCODERS: xml, json
* USED COMMON CONTROL STRUCTURES: if, for range, error check, and etc
```
### Screenshot
![Screenshot PathFinder](https://lh3.googleusercontent.com/R_X7_FtRH2Zu4JZcL_nYPNnn5Zhh1441Sa_dBr6vkAekr97DRxMxfx30J9szQN2NZP-3WaFwh2wI74rosxSRnh7Glr5agb-n0O9Hu_Xd8P72q61CBep3kGVjUNuD88EYEXkWnwegWPyf8rRjUEqxeUpcKTCW22MqUaOxQpFkZBFkd5jK8P3O4bKDjslZ9ABtDZGiCAj6GNu_I3AjDr6OQfsB69u6c6dnVwC3WD7tORUVwJYHqK9LYi_84IkcrBZgOZRLWnsgXqZEDXvBHbpvnzIG8ihXlJeBZSju034CloMPL2tt8EI7W1-KKjAPNpjFAHTEE769AFtp4nlTbf13PNv6Dx4PLrAdrmsxOykzivmQ2cUnosI__9NmsLUv4Fr1R-ZhgAmOju6BSSrhOE2rhV8esyuJeMV9HIoRt0z0JStl3Ui_EI0hwKWpkB-Ly580Ff4le5ct9aq8Z2ptAio1yx73HefBrdb8NtJwRtUj7xCenXjj-3i9MGlEcZU_smM8pMU5C-V7oXt-ajjP-3UKAcmXz5085GzuxWVncB87wWwjOp6sr448h1PJJWBxNVLIrGs=w1357-h922-no "screenshot PathFinder") 

