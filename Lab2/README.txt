Integrantes Grupo 16:
    José Castro, 202073550-6
    Florencia Ramírez, 202073522-0

☆.。.:*・°☆.。.:*・°☆.。.:*・°☆.。.:*・°☆.。.:*・°☆

Instrucciones de compilación:
    Esta tarea fue hecha con go 1.22.1. 

    Se deben abrir 2 terminales distintas, en la primera terminal se ejecutará el código en Docker del servidor central.
    Para la creación del contenedor.
        docker build -t main-server -f DockerfileServer .
    Para la ejecución del contenedor.
        docker run -d --name main-server-container -p 8080:8080 main-server
    Finalmente para ver los prints por consola.
        docker logs -f --tail 10 main-server-container
    
    Luego, en la segunda terminal se ejecutará el código de los capitanes.
        go run Captains/Captains.go

☆.。.:*・°☆.。.:*・°☆.。.:*・°☆.。.:*・°☆.。.:*・°☆

Informe:
    Mientras que esto podría ser una solución implementada para cuando desean, por ejemplo, cubrir distancias demasiado 
    grandes tales que una sola central sería capaz de cubrirla, esto complejisaría el problema en términos de comunicación 
    síncrona y consistencia, ya que sería pertinente que ambos "servidores" siempre tuvieran la misma información para 
    evitar errores o inconsistencias en los números de tesoros.

☆.。.:*・°☆.。.:*・°☆.。.:*・°☆.。.:*・°☆.。.:*・°☆

Comentarios extra:
    - El código no termina la ejecución, para terminar las ejecuciones se debe presionar Ctrl + C en ambas terminales.