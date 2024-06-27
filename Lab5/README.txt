Integrantes Grupo 16:
    José Castro, 202073550-6
    Florencia Ramírez, 202073522-0

☆.。.:*・°☆.。.:*・°☆.。.:*・°☆.。.:*・°☆.。.:*・°☆

Instrucciones de compilación:
    Esta tarea fue hecha con go 1.22.1.
    La distribución de las máquinas virtuales es:
        - Máquina 1 (061): Comandante Kais, Servidor Fulcrum 1.
        - Máquina 2 (062): Ingeniero Jeth, Servidor Fulcrum 2.
        - Máquina 3 (063): Ingeniero Malkor, Servidor Fulcrum 3.
        - Máquina 4 (064): Broker.
    Se deben abrir 7 terminales cada una conectada a una máquina dependiendo de su entidad y luego ir a la carpeta Grupo-16-Lab5/Lab5.
    Ejecutar dentro de las máquinas correspondientes los códigos de las entidades en el siguiente orden:
        $ make broker
        $ make server1
        $ make server2
        $ make server3
        $ make jeth
        $ make malkor
        $ make comandante

☆.。.:*・°☆.。.:*・°☆.。.:*・°☆.。.:*・°☆.。.:*・°☆

Comentarios extras:
    - Se muestra la ejecución con make aunque no está instalado en las máquinas, pero corresponden a los comandos que se deberían usar para compilar cada entidad.
        - Se pide la contraseña de la máquina al ejecutar los comandos ya que no se puede hacer sin el modo sudo, pero estas no se incluyeron en el archivo.
    - Borramos los contenedores y las imágenes de las máquinas después de nuestras pruebas pero si da problemas porque ya existen se deben borrar ejecutando:
        $ sudo docker image prune
        $ sudo docker container prune
    - No se implementó por completo la consistencia entre los servidores, funciona relativamente bien pero intenten no ejecutar muchos comandos antes que se propaguen :<
        - Tampoco se implementaron los modelos de consistencias en los ingenieros y el comandante.