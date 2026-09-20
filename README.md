# OnThePitch

*Proyecto realizado por Jaime Rojas Herrera*

### Descripción del Problema

A lo largo de estos años, me he dado cuenta de que los campos de fútbol de mi pueblo están casi siempre vacíos. Siento que los aficionados de este deporte no se animan a practicarlo, debido seguramente a la dificultad de encontrar otros con los que jugar. Incluso a mí hay veces que me ha pasado lo siguiente: comentas por tú grupo de amigos para ir y echar un rato jugando, pero ninguno acaba pudiendo ir, quedándote con las ganas. Además, en mi pueblo no se celebran muchos torneos locales. Para terminar, otro problema con el que me he encontrado es con aficionados de un equipo que no asisten a los estadios porque no tienen a nadie con quién ir. En estos tiempos, con las redes sociales a la orden del día, me sorprende la carencia de un medio para poder conectar con amantes de uno de los deportes con más renombre en la actualidad.

### Juego de Rol

[Fotografía de la tarjeta de rol Desarrollador](Desarrollador.jpeg)
[Fotografía de la tarjeta de rol Cliente](Cliente.jpeg)
[Fotografía de la tarjeta de rol Validación](Validacion.jpeg)

### Configuración Previa Realizada

[Captura de la configuración de SSH](ssh.png)
[Captura de la configuración de GIT con nombre y email](config.png)

### Lista de Comprobación ¿Cumple el problema con lo requerido?

#### ¿Se trata de un problema real del que se tenga conocimiento personal?

Si, de hecho, en mi propio pueblo tengo constancia de varias personas que no encuentran nadie con quién jugar.

#### ¿Se trata de un problema que para solucionar requiera el despliegue de una aplicación en la nube?

Correcto, es totalmente necesario el uso de una aplicación/plataforma conectada en la red para que los usuarios puedan solicitar personas con las que poder jugar. Un usuario debería acceder al servicio, introducir su nivel y mandar un aviso por la plataforma indicando que busca gente con la practicar el deporte, aviso que posteriormente se enviará a usuarios cerca de su zona. Sin estar la aplicación desplegada en la nube, esto sería imposible.

#### ¿La solución requiere una cierta cantidad de lógica de negocio, en vez de solucionarse sólo almacenando y buscando?

Efectivamente. Esta solución implica que los usuarios tienen un nivel (que en un principio ellos mismos eligen), por lo que a la hora de realizar los emparejamientos, es necesario que haya un equilibrio para que los equipos no acaben descompensados. Además, los propios usuarios, siempre y cuando un partido haya llegado a su conclusión, pueden darle una valoración al partido que haya hecho un usuario en particular, lo cual se tendrá en cuenta para ajustar su nivel en posteriores partidos. Por último, jugadores con un nivel avanzado claramente no van a emparejarse con jugadores de nivel bajo, por lo que en el momento de mandar avisos se descartarían jugadores con flébil nivel.

#### ¿Se ha incluido la configuración del repositorio y se ha enlazado desde el `README`?

Si, por supuesto. La configuración se encuentra arriba. También se ha enlazado con las fotografías del juego de rol.

#### ¿El estudiante tiene todos los datos necesarios para poder resolver el problema, o va a requerir que el usuario los introduzca?

Es necesario que el usuario introduzca su nivel, o, a medida que vaya jugando, que otros usuarios determinen su nivel en base a valoraciones. Eso sí, datos referentes a estadios (para que los usuarios puedan encontrar personas con las que asistir) se extrarían directamente de Google Maps. Más allá de eso, es una aplicación por y para los usuarios, por lo que ellos mismos son los datos con los que trabajar.



