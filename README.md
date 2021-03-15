# Deputy
Implementation of the code test for Deputy.

## Prerequisites
Requires golang be installed

## Installation
To install the executable, execute

```
git clone https://github.com/gilmae/deputy.git
cd deputy
go install
```

## Testing 
To run the test suite, from within the deputy folder, execute:

`go test ...`

## Execution
Assuming the source was installed as above, the deputy executable can be run with:

`deputy --data data_file --supervisor user_id`

To execute from source directly, run:

`go run main.go --data data_file --supervisor user_id`

data_file will default to `./sample.json`
user_id will default to `1`

### Data File
The executable requires a data file that defines roles and users. This data file is a json file with teh format:

```
{
    "roles": [
        {
            "Id": 1,
            "Name": "Sample role",
            "Parent": 0
        },
        ...
    ],
    "users": [
        {
            "Id": 1,
            "Name": "Sample user",
            "Role" 1
        },
        ...
    ]
}
```

A sample data file has been included in the source, `sample.json`