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
            "id", "in"                                                   // What should be compared 
        ],
        "level": "warn",                                                 // default log-level for Test
        "datas": [
            {
                "level": "debug",                                        // log-level for data
                "in":   {
                    "key": "string",
                    "value": "InValue"
                },
                "params": [{"key": "id", "value": "ParamIdValue"}],
                "out": {}                                                // values to compare
                "exp": {                                                 // expected values
                    "id": "ParamIdValue",
                    "in": "InValue"
                }
            }
        ]
    }
]   
## Compare Values
TestDataChecker checks out and exp for keys from config read in by GetTestDataConfig 
## Sample how to use
One use is shown in the file testtextdatasimple_test.go
