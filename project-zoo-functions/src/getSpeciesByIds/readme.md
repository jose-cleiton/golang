### Bem-vindo à função getSpeciesByIds!

Esta função é responsável por buscar as espécies de animais a partir de seus ids.
Ela pode receber um ou mais ids como parâmetros e retorna um array contendo as espécies referentes a esses ids.

Exemplo de uso:
```bash
id1 := "0938aa23-f153-4937-9f88-4858b24d6bce"
id2 := "ef3778eb-2844-4c7c-b66c-f432073e1c6b"

species, err := getSpeciesByIds(id1, id2)
if err != nil {
  fmt.Println(err)
} else {
  fmt.Println("species:", species)
}
```
## Observações técnicas:


* O parâmetro da função pode ser alterado para atender ao requisito proposto.

* A função deve receber um ou mais ids como parâmetros.
* A função lê os dados de uma fonte externa (arquivo JSON), portanto é necessário garantir que o arquivo esteja presente e estruturado corretamente.
  








