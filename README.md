# ClientGpt



# GPT-3 Chat API Client

Este paquete proporciona estructuras y funciones para interactuar con la API de chat de OpenAI GPT-3.

## Estructuras


### ChatCompletionRequest

Esta estructura representa una solicitud de chat para la API de OpenAI GPT-3.

```go
type ChatCompletionRequest struct {
	Model    string        `json:"model"`
	Messages []ChatMessage `json:"messages"`
}
```
## ChatMessage

Esta estructura representa un mensaje en una conversación de chat.

```go
type ChatMessage struct {
	Role    string `json:"role"`
	Content string `json:"content"`
}
```


# Funciones

## ReqChatGpt

Esta función realiza una solicitud de chat a la API de OpenAI GPT-3 y devuelve un objeto ChatCompletion que contiene la respuesta.

```go
func (chat ChatCompletionRequest) ReqChatGpt(contet, apiKey string) (model.ChatCompletion, error)

```

## ReadCompressedBody
Esta función lee el cuerpo comprimido de una respuesta HTTP y devuelve los bytes descomprimidos.

```go

func ReadCompressedBody(resp *http.Response) ([]byte, error)

```


# UngzipReader
Esta función crea un lector para descomprimir el cuerpo de una respuesta HTTP si está comprimido en formato gzip.

```go

func UngzipReader(resp *http.Response) (io.ReadCloser, error)

```

# Ejemplo

```go
package main

import (
	"github.com/sebarray/ClientGpt/service"
)

func main() {
	var client service.ChatCompletionRequest
	respose, err := client.ReqChatGpt("Hello, how are you?", "apikey")
	if err != nil {
		panic(err)
	}
	println(respose)
}
```