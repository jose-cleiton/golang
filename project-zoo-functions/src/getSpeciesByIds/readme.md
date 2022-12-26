# GetSpeciesByIds

Esta função é responsável pela busca das espécies de animais por id. Ela retorna um array contendo as espécies referentes aos ids passados como parâmetro, podendo receber um ou mais ids.

## Parâmetros

- `zoo`: um struct do tipo `Zoo` que representa o zoológico.
- `ids`: um array de strings contendo os IDs das espécies a serem buscadas.

## Retorno

Um array de structs do tipo `Species`, contendo as espécies referentes aos IDs fornecidos.

## Exemplos

```go
import "zoo"

// Busque a espécie de leão pelo ID
species := zoo.GetSpeciesByIds(zoo, "0938aa23-f153-4937-9f88-4858b24d6bce")

// Busque as espécies de leão e tigre pelos IDs
species := zoo.GetSpeciesByIds(zoo, "0938aa23-f153-4937-9f88-4858b24d6bce", "e8481c1d-42ea-4610-8e11-1752cfc05a46")
