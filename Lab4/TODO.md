# Enorme lista de cosas que hacer pal lab （＞人＜；）

> [!NOTE]
> En vola se me olvidó algo pero creo que es eso nomás, peor caso lo actualizo y te aviso :>

## Mercenarios

### Solo por consola
- [ ] Opción para ver el monto acumulado
- [ ] Enviar estado al director
- [x] Pedir decision por consola

### Bots
- [x] Decisiones al azar

### Todos
- [ ] Enviar decisiones a director
- [ ] Obtener respuesta si sobrevive o no el piso
- [ ] Esperar para el inicio de piso

## Director
- [ ] Interfaz por consola
    - [ ] Opción para pedir ver todas las decisiones de un mercenario
    - [ ] Opción para comenzar el siguiente piso
        - Tiene que esperar a todos los mercenarios para iniciar
- [ ] Lógica de cada piso
    - [ ] Piso 1 
        - Obtener X e Y para obtener los rangos de las armas, opción A va de 0 a X, opción B va de X a Y, y opción C va de Y a 100
        - Por cada mercenario genera un número al azar del 0 al 100
        - Si coincide en el rango del arma, sobrevive. Sino, muere
    - [ ] Piso 2
        - Selecciona al azar entre A y B (rand.Intn(2))
        - Sobreviven los mercenarios que coinciden con el pasillo seleccionado
    - [ ] Piso 3
        - Selecciona al azar 5 números del 1 al 15
        - Por cada mercenario, el número i del mercenario coincide con el número i del director al menos en 2 rondas
        - Ej: mercenario [1,4,9,2,11] y director [1,9,3,2,15] coinciden el primero (1) y el cuarto (2) -> 99% segura que es así y no que no importe el orden y son solo los números en común que tienen (siendo 1, 2 y 9)
- [ ] Mostrar por consola cuando muere un mercenario
    - [ ] Mandar mensaje al DoshBank con el nombre del mercenario que murió
    - [ ] Informar al mercenario que murió para terminar su ejecución
- [ ] Mostrar por consola los mercenarios vivos al final del piso
    - Si queda solo un mercenario vivo al final de un piso, se termina y considera que perdió
- [ ] Mostrar por consola los mercenarios que ganaron (cuando terminaron los 3 pisos si quedan mercenarios vivos)
- [ ] Enviar las decisiones de los mercenarios al NameNode.
    - Indica el nombre del mercenario, el piso actual, y la(s) decision(es) que tomó

## DoshBank
- [x] Crear un archivo txt registrando las muertes de los mercenarios y el monto actual
- [ ] Actualizar el archivo txt cada vez que se notifica la muerte de un mercenario
    - Aumenta el monto actual en 100M (100000000)
- [ ] Enviar el monto acumulado al director cuando lo pida

## NameNode
- [x] Crear un archivo txt con la información de las decisiones de todos los mercenarios
- [ ] Escoger al azar en que DataNode se va a guardar la(s) decision(es) del mercenario en un piso específico
- [ ] Actualizar el archivo txt con la información del mercenario, piso e IP del DataNode en que se guarda la información
- [ ] Obtener y enviar todas las decisiones de un mercenario al director
    - [ ] Pedir al DataNode correspondiente la(s) decision(es) de un mercenario y piso específico
    - [ ] Juntar todas las decisiones para luego enviarlas al director

## DataNodes
- [x] Crear un archivo txt para el mercenario y piso correspondiente guardando la(s) decision(es) que tomó
- [ ] Obtener la(s) decision(es) de un mercenario y enviarla(s) al NameNode cuando lo pida

## Servicio - Mensajes gRPC
- [ ] Petición para informar estado (mercenario - director)
- [ ] Enviar decisiones por piso (mercenario - director)
- [ ] Pedir ver monto acumulado (mercenario - director)
- [ ] Informar resultado de piso (director - mercenario)
- [ ] Enviar monto acumulado (director - mercenario)
- [ ] Informar inicio de piso (director - mercenario)
---
- [ ] Pedir monto acumulado (director - doshbank)
- [ ] Enviar monto acumulado (doshbank - director)
---
- [ ] Enviar decisiones de mercenarios (director - namenode)
- [ ] Pedir todas las decisiones de un mercenario (director - namenode)
- [ ] Enviar todas las decisiones de un mercenario (namenode - director)
---
- [ ] Enviar decision para registrar (namenode - datanode)
- [ ] Pedir decision (namenode - datanode)
- [ ] Enviar registro de decisiones (datanode - namenode)

## Mensaje RabbitMQ
- [ ] Informar muerte de mercenario (director - doshbank)

## Información de máquinas
- Director, DoshBank y NameNode en máquinas distintas
- Cada DataNode debe estar en una máquina distinta
- NameNode no puede ejecutarse en la misma máquina que un DataNode
- Técnicamente los mercenarios pueden ejecutarse en máquinas distintas pero yo pensaba hacerlo usando lo del sync
- Idea posible de máquinas:
    - Máquina 1 (061): Director y DataNode1
    - Máquina 2 (062): DoshBank y DataNode2
    - Máquina 3 (063): NameNode y Cola RabbitMQ (esto como que lo saqué del lab de la Sofi, no se si tiene que ser un archivo aparte)
    - Máquina 4 (064): DataNode3 y Mercenarios.