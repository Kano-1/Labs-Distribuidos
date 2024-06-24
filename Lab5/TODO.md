# Cosas que hacer (•ˋ⌒ˊ•)

## Servidores Fulcrum
- [x] Almacenar registro de los soldados enemigos en un sector en un archivo txt
    - Escribe una línea por cada base en el sector, ej: "SectorAlpha Campamento1 15"
    - Si no hay un archivo para el sector se tiene que crear
- [x] Actualizar el reloj de vector con la cantidad de cambios hechos
- [x] Tener un log de registros de los cambios que se han hecho
    - Ej: AgregarBase SectorAlpha Campamento2 13
    - Cuando se propagaron los cambios entre los servidores tienen que ser **borrados y creados** de nuevo ?
- Técnicamente si está implementado todo eso, falta el que reciba la información con gRPC
- [ ] Propagar los cambios cada 30 segundos
    - [ ] Realizar ***Merge*** si hay inconsistencias 

## Broker
- [x] Escoger un servidor al azar para el ingeniero haga sus cambios en él
- [ ] Escoger un servidor al azar para el comandante cuando quiera consultar información
- [ ] Manejar las inconsistencias
    - Si el ingeniero detecta una, el broker debe decidir cuál es el mejor servidor que puede ayudar al ingeniero ?

## Ingenieros
- [x] Preguntar por consola los cambios en los servidores
    - [x] Agregar una base a un sector
    - [x] Renombrar una base de un sector
    - [x] Actualizar el valor de enemigos en una base
    - [x] Borrar una base de un sector
- [ ] Utilizar ***Read your writes*** para llevar a cabo la consistencia
    - Pueden guardar información en memoria
    - Revisan si hay inconsistencias con esto
    - Tiene que leer después que escribe ??

## Comandante
- [x] Consultar por consola la información sobre un sector y base
    - [x] Obtener la cantidad de enemigos en una base
        - Incluye el reloj de vector del servidor
- [ ] Utilizar ***Monotonic reads*** para mantener la consistencia
    - Pueden guardar información en memoria
    - También revisa si hay inconsistencias

## Consistencia eventual - Merge
- [ ] Definir un nodo dominante, el que esté mas actualizado
- [ ] Comparar las operaciones con los logs de los servidores para aplicar los cambios del nodo dominante
- [ ] Actualizar los relojes de vectores
    - Ej: [1, 0, 0], [0, 1, 0] y [0, 0, 0] donde todos pasan a ser [1, 1, 0]
- [ ] Utilizar método de resolución libre para casos particulares ?
    - [ ] Especificar en el README
- El objetivo es que al finalizar el proceso no existan diferencias entre los 3 servidores
    - Como es ***consitencia eventual*** da lo mismo que se "pierda" información

## Mensajes gRPC
**No sé si eso sería todo y como funciona bien lo de las inconsistencias :<**
- [x] Petición escribir información (ingeniero - broker)
- [x] Escribir información (ingeniero - servidor)
- [ ] Informar inconsistencia (ingeniero - broker)
- [x] Entregar dirección servidor (broker - ingeniero)
---
- [ ] Informar inconsistencia (servidor - broker) ?
- [ ] Propagar cambios (servidor - servidor)
---
- [x] Petición obtener información (comandante - broker)
- [x] Consultar información (comandante - servidor)
- [x] Entregar dirección servidor (broker - comandante)
- [x] Entregar información (servidor - comandante)

## Separación máquinas
- 3 servidores en máquinas distintas
- Broker en una máquina distinta a los servidores
- Comandante, dos ingenieros y broker en máquinas distintas
- Posible distribución:
    - Máquina 1 (061): Comandante Kais, Servidor Fulcrum 1
    - Máquina 2 (062): Ingeniero Jeth, Servidor Fulcrum 2
    - Máquina 3 (063): Ingeniero Malkor, Servidor Fulcrum 3
    - Máquina 4 (064): Broker