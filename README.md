#Gauntlet Implementation
Design Details: 

add all adindex id to a trie using the properties
search all available ads using the user properties

for example:
userprofile {
    "age": 24,
    "gender" : "man",
    "country" : "US",
    "language" : ["cn","en"],
    "appIds" : ["robinhood.com", "postmate.com"]
}

ads [
    {
        "age": ">18&<35",
        "gender" : "man",
        "country" : "US",
        "language" : ["en"],
        "appIds" : ["robinhood.com"]
    },
    {
        "age": ">45",
        "gender" : "man",
        "country" : "CN",
        "language" : ["en"],
        "appIds" : ["robinhood.com"]
    }
    ,
    {
        "age": ">45",
        "gender" : "woman",
        "country" : "CN",
        "language" : ["en"],
        "appIds" : ["wechat.com"]
    }
]

TreeNode{
    type: ["string", "range", "list", "exclude_list"]
    rule: {
        "string":"match",
        "range":"compare",
        "list":"in list",
        "exclude_list": "not in list",
    },
    value : original value
}

age : {type : string, rule : {greater, 18} }
