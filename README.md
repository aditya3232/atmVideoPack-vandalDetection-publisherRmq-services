
# API Spec Publisher AtmVideoPack Vandal Detection

## 1.1 Vandal Detection

### 1.1.1 POST :: Vandal Detection

Request :
- Method : POST
- Endpoint : `localhost:3434/publisher/atmvideopack/v1/vandaldetection/create`
- Header :
    - Content-Type : application/x-www-form-urlencoded
    - x-api-key : required
- Body (form-data: x-www-form-urlencoded) :
    - tid : string, required
    - date_time : string, required
    - person : string, required
    - file : file, required
- Response :

```json 
{
    "meta": {
        "message": "string",
        "code": "integer",
    },
        "data":{
            "tid": "string",
            "date_time": "string",
            "person": "string",  
            "converted_file": "string"
        }
 }
```






