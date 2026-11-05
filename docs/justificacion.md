# Gestor de Dependencias

El gestor de dependencias es clave para poder gestionar las dependencias externas que se van a usar en el repositorio y que otros colaboradores deberán tener para poder trabajar correctamente con el código. Los criterios para decidir un buen gestor de dependencias son:

- Se sigue el estándar del lenguaje, en este caso Go.
- Se siguen las mejores prácticas.
- Tratar de evitar la deuda técnica en la medida de lo posible.
- La instalación del gestor sea lo más simple y rápida posible.
- El gestor funcione bien con las dependencias que se van a utilizar en el proyecto.

Teniendo en cuenta estos aspectos, en primer lugar se ha decidido buscar el estándar de Go para gestionar las dependencias. Indagando en la documentación oficial de Go, el estándar utilizado para gestionar las dependencias es Go Modules, que viene integrado junto a la instalación de Go (https://go.dev/doc/modules/managing-dependencies). Es por tanto, la mejor práctica para la gestión de las dependencias, pues el gestor oficial que Go insta a utilizar en su página de documentación. Utilizar el gestor integrado y totalmente estable permite reducir la deuda técnica ostensiblemente, pues utilizar un gestor de dependencias obsoleto implica tener que cambiarlo con posterioridad. La instalación del lenguaje Go instala a su vez el gestor de dependencias Go Modules, por lo que la instalación es tan simple como instalar el propio lenguaje. Aún no se saben con exactitud las dependencias que se van a usar en el proyecto (relacionadas con OSM), pero Go Modules es un gestor que funciona de forma fluida e integrada con Go.

Se ha buscado a su vez información sobre otros gestores de dependencias que no sean el estándar, sin embargo todos están obsoletos y recomiendan utilizar Go Modules:

- Godep: La página del repositorio indica que la herramienta ha sido archivada y que se recomienda utilizar dep: ![Captura del repositorio de Godep](IMAGENES/justi1.png)

- gb: Era más bien una herramienta de build alternativa. El repositorio fue archivado el 6 de Marzo de 2020 y ahora es solo de lectura: ![Captura del repositorio de Godep](IMAGENES/justi2.png)

- glide: Se utilizaba mucho en su día, pero al igual que las demás dejo de recibir mantenimiento. En su repositorio recomiendan migrar: ![Captura del repositorio de Godep](IMAGENES/justi3.png)

- dep: Era más bien un experimento y en la propia documentación oficial indican que lo suyo es utilizar un prototipo alternativo (Go Modules): ![Captura del repositorio de Godep](IMAGENES/justi4.png)

Teniendo en cuenta que todas han quedado obsoletas frente a Go Modules, se ha tomado la decisión de utilizar el gestor de dependencias integrado para que todo fluya de la mejor manera.



