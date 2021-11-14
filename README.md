# testtextdatasimple
To make testing with text data simple
## configuration
The configuration is read from a json config file
### Syntax of the json config file
[
    {   "name": "TestFuncAlias", "alias": "TestFuncId"},
    {   "name": "TestFuncNoAlias", "alias": "", "level": "debug"},
    {
        "name": "TestFuncId",
        "keys": [
            "id", "in"
        ],
        "level": "warn",
        "datas": [
            {
                "level": "debug",
                "in":   {
                    "key": "string",
                    "value": "InValue"
                },
                "params": [{"key": "id", "value": "ParamIdValue"}],
                "est": {
                    "id": "ParamIdValue",
                    "in": "InValue"
                }
            }
        ]
    }
]    
