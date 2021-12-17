# URLshortener-golang

"Encurtador" de URLs em golang: uma string ```https://go.io/xxxxxx``` é gerada a partir de uma URL. O código ```xxxxxx``` é alfanumérico e único.

# POST
> localhost:8000/send/**URL**

**Retorno:** JSON com dados sobre a URL {ID, data e hora da operação, tempo de processamento, URL original e URL encurtada}

# GET
> localhost:8000/retrieve/**URL encurtada**

**Retorno:** JSON com dados sobre a URL {ID, data e hora da operação, tempo de processamento, URL original e URL encurtada}
