# mutant-detector
Servicio para detección de mutantes.

## Consideraciones

Una idea que se tuvo en cuenta para el desarrollo es que sea facilmente aplicable a buscar secuencias de tamaño N en una matriz con K digitos distintos, para esto se tomaron dos decisiones:
- Que la longitud de las secuencias a buscar pueda ser configurable
- Que el algoritmo de busqueda no tenga ninguna restriccion en cuanto a los caracteres que recibe (La validacion de que sea un DNA valido esta separada y es facilmente modificable)

## Instalación

```
cd $GO_HOME/src
git clone https://github.com/mnoliveira/mutant-detector.git
```
## Dependencias
```
go get github.com/gin-gonic/gin
go get go.mongodb.org/mongo-driver/mongo
go get github.com/olebedev/config
```

## Compilación
```
go build
```

## Ejecutar
```
go run mutant-detector
```

## Host
Para hostear el servicio se utilizaron los servicios de AWS. En particular:
* **EC2** - para la instancia de MongoDB
* **Elastic Beanstalk** - para el servicio*

\* Por temas de costo, no se habilito el autoscaling.

## APIs  
URL Host: http://mutant-env.eba-zgsi5a4x.us-east-1.elasticbeanstalk.com/ \
URL Local: http://localhost:5000

#### POST /mutant
Devuelve si el dna corresponde a un mutante.
##### Request:
```json  
{
    "dna":["ATGCGA","CAGTGC","TTATGT","AGAAGG","CCCCTA","TCACTG"]
}
```  
##### Response:
###### En el caso de ser mutante:
Status: 200
```json  
{
    "message": "Mutante detectado!"
}
```
###### En el caso de ser humano:
Status: 403
```json  
{
    "error": "Es solo otro humano."
}
```
###### Otros posibles errores:
Status: codigo del error
```json  
{
    "error": "Motivo del error"
}
```

#### GET /stats
Devuelve el reporte de detección de mutantes.
##### Response:
```json  
{
    "count_mutant_dna": 2,
    "count_human_dna": 3,
    "ratio": 1.5
}
```  


## Test
Para la ejecucion de los test se utilizaron las herramientas provistas por el IDE Jetbrain Goland. Donde se obtuvo un coverage de 86.8%.

![coverage](docs/coverage.png)
