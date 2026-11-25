# Gestor de Dependencias

Teniendo en cuenta estos aspectos, en primer lugar se ha decidido buscar el estándar de Go para gestionar las dependencias. Indagando en la documentación oficial de Go, el estándar utilizado para gestionar las dependencias es Go Modules, que viene integrado junto a la instalación de Go (https://go.dev/doc/modules/managing-dependencies). Se ha buscado a su vez información sobre otros gestores de dependencias que no sean el estándar, sin embargo todos están obsoletos y recomiendan utilizar Go Modules. Teniendo en cuenta esto, se ha tomado la decisión de utilizar el gestor de dependencias integrado para que todo fluya de la mejor manera.

# Gestor de Tareas

La decisión sobre el gestor de tareas depende de los siguientes criterios:

- Tiempo de instalación.
- Tiempo de ejecución.

- make: permite automatizar comandos repetitivos. Su instalación en muchos entornos no es necesaria (viene instalado por defecto). Es la más simple de configurar aunque su sintaxis en ocasiones puede resultar algo liosa. Buscando acerca de este gestor de tareas, encontré un post en Reddit dónde cataloga este gestor como uno de los mejores para este tipo de tareas (https://www.reddit.com/r/golang/comments/5f0st7/preferred_task_manager_for_go_projects/). Se ha hecho una medición del tiempo de instalación en un docker con la imagen de Ubuntu, que no viene por defecto instalado, dando el siguiente resultado:
![Captura de la medición de make](../IMAGENES/justi5.png)
Otro criterio importante a la hora de decantarse por un gestor de tareas u otro es el tiempo de ejecución de las tareas. En el caso de make, la orden check tarda lo siguiente:
![Captura del tiempo de ejecución de make](../IMAGENES/justi8.png)

- Task: es una herramienta que la propia documentación de Go trata como un gestor de tareas con el objetivo de ser más simple y sencillo de usar que make (https://pkg.go.dev/github.com/go-task/task). El mayor inconveniente con respecto a make es el proceso de instalación, que es algo más complejo:
![Captura de la medición de Task](../IMAGENES/justi6.png)
El binario ocupa 74 MB (mucho más que make) y hay que utilizar estos comandos para añadir /go/bin al PATH:
echo 'export PATH=$PATH:$HOME/go/bin' >> ~/.bashrc
source ~/.bashrc
En cuanto al tiempo de ejecución, Task tarda lo siguiente:
![Captura de la medición de Task](../IMAGENES/justi10.png)

- Just: gestor de tareas estándar de Rust. Al igual que Task, tiene una sintaxis más clara que make, sin embargo, es un binario externo escrito con Rust, por lo que los entornos o colaboradores del proyecto lo tienen que instalar por separado.
![Captura de la medición de Just](../IMAGENES/justi7.png)
El tiempo de instalación también es algo más elevado que make y Task.
En cuanto al tiempo de ejecución, Just tarda lo siguiente:
![Captura del tiempo de ejecución de Just](../IMAGENES/justi9.png)

- tj/robo: este gestor de tareas tiene una sintaxis nativa en YAML y está diseñado exclusivamente a tareas. El problema principal es que su mantenimiento es escaso a día de hoy. La instalación tampoco viene integrada como make.
![Captura de la medición de tj/robo](../IMAGENES/justi11.png)
En cuanto al tiempo de ejecución, tj/robo tarda lo siguiente:
![Captura del tiempo de ejecución de tj/robo](../IMAGENES/justi12.png)

Aunque todas las opciones son aceptables, siendo todas estables y mantenidas, el hecho de tener que depender de binarios externos tanto para Task como para Just, hace que me decante por utilizar make, pese a tener una sintaxis algo más quisquillosa. Además, teniendo en cuenta el importante criterio sobre el tiempo de ejecución, task tarda bastante más que make y just, cuyos ambos gestores de tareas tienen tiempos similares, ligeramente inferiores en make, por lo que otra razón para decidirse por este simple y efectivo gestor. Se descarta tj/robo pues es un proyecto algo antiguo.

# Herramientas para Tests

## Biblioteca de Aserciones

Para elegir la biblioteca de aserciones, se siguen los siguientes criterios:

- Dependencias. Si la biblioteca supone añadir demasiadas dependencias, esto impacta en el rendimiento del proyecto y puede ser un problema

- Tiempo de Instalación.

Opciones:

- "Simular" aserciones. Testing, el framework nativo que usa Go para escribir los tests, no dispone como tal de una biblioteca de aserciones, aunque es posible hacer las comparacions de forma manual, es decir, la aserción sería un if que compara y llama a "t.Errorf" si no se cumple. No requiere dependencias y no tiene tiempo de instalación pues no es necesario instalar nada más.

- testify. Funciona tanto como para escribir los tests como para usarse como biblioteca de aserciones. Se instala con "go get" y esta compuesta de varios paquetes, siendo "assert" el paquete que incluye las aserciones. Requiere de dependencias externas y el tiempo que tarda en instalarse es el siguiente:
![Captura del tiempo de instalación de testify](../IMAGENES/justi13.png)

- Gomega. Gomega es una Biblioteca de "matchers" que viene emparejada con Ginkgo (Test Runner) dónde las aserciones se pueden leer como frases (Expect(x).To(Matcher)). Requiere de dependecias y el tiempo que tarda en instalarse es el siguiente:
![Captura del tiempo de instalación de Gomega](../IMAGENES/justi14.png)

La decisión final es utilizar testify para escribir las comparaciones y prescindir de una biblioteca adicional que necesite de dependencias y un tiempo de instalación. En Gomega por ejemplo es de 2 segundos, algo de lo que se puede prescindir perfectamente para no comprometer el rendimiento.

## Herramientas CLI

Se decide utilizar el estándar de Go, "go test", el cúal además es compatible con todos los frameworks que se van a ver a continuación. Se han visto otras alternativas que aportan sobre "go test", como gotestsum que devuelve una salida legible, ginkgo watch que da Feedback rápido al desarrollar, ginkgo CLI que reparte los procesos paralelamente y gocov o go-carpet que devuelve informes más detallados sobre el testeo, Sin embargo, el requerir de tiempo de instalación y de tener que gestionar dependencias externas hace que el valor aportado no merezca la pena, por lo que se acaba optando por "go test".

## Test Runner

Para elegir el Test Runner, se siguen los siguientes criterios:

- Dependencias.

- Tiempo de Instalación.

- Testing. Framework nativo y estándar de Go. No requiere de tiempo de instalación adicional. Se usa el comando "go test" para ejecutar los ficheros con "_test.go" que incorporan los tests sobre un paquete en específico.

- Ginkgo. Es un Test Runner que habría que instalar adicionalmente y trae también varias dependencias consigo. Es un framework BDD (desarrollo guiado por comportamiento) completo. El tiempo de instalación es el siguiente:
![Captura del tiempo de instalación de Ginkgo](../IMAGENES/justi15.png)

- franela/goblin. Framework minimalista, inspirado en Mocha, que añade sobre Testing. El tiempo de instalación es el siguiente:
![Captura del tiempo de instalación de franela/goblin](../IMAGENES/justi16.png)

Se decide finalmente utilizar testing como el Test Runner, es el estándar de Go, no necesita instalar nada por lo que el tiempo de instalación es nulo. Como nota adicional "go test" no escribe en el directorio de fuentes, tan solo en una caché fuera del repositorio (GOCACHE)








