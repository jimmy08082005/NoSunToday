# HealthyRoutes

*Proyecto realizado por Jaime Rojas Herrera*

### Descripción del Problema

Las personas mayores, con movilidad reducida o con algún problema cardiovascular sufren cada día moviéndose por las calles de Granada. Cuando necesitan ir a un lugar desconocido, hacen uso de Google Maps, pensando que la ruta designada será óptima. Y la verdad, probablemente sea la mejor posible, pero no para ellos. Altos desniveles, largas callejuelas sin una pizca de sombra en verano, horas y horas de camino sin bancos o fuentes en las que poder tomar un respiro... En definitiva, Google Maps no se ajusta a estas condiciones adversas que provocan múltiples quebraderos de cabeza para estas minorías desfavorecidas.

### Juego de Rol

- [Fotografía de la tarjeta de rol Desarrollador](Desarrollador.jpeg)
- [Fotografía de la tarjeta de rol Cliente](Cliente.jpeg)
- [Fotografía de la tarjeta de rol Validación](Validacion.jpeg)

### Configuración Previa Realizada

- [Captura de la configuración de SSH](ssh.png)
- [Captura de la configuración de GIT con nombre y email](config.png)

### Lista de Comprobación ¿Cumple el problema con lo requerido?

#### ¿Se trata de un problema real del que se tenga conocimiento personal?

Si, este verano se han reportado numerosos golpes de calor que quizás se podrían haber sofocado si siguiesen una ruta ajustada para ellos.

#### ¿Se trata de un problema que para solucionar requiera el despliegue de una aplicación en la nube?

Correcto, es necesario recabar en cada momento condicionantes como la temperatura, la posición del sol, obstaculos en la ruta... Todo esto en tiepo real, para poder proporcionar al usuario la mejor ruta posible. De no estar en la nube, actualizar estos datos en tiempo real no sería posible.

#### ¿La solución requiere una cierta cantidad de lógica de negocio, en vez de solucionarse sólo almacenando y buscando?

Efectivamente. En primer lugar se extraen los datos de plataformas como el ayuntamiento de granada, de plataformas de meteorología, incluso del propio Google Maps. Posteriormente estos datos serán analizados y se calcularán, por ejemplo, las zonas de sombra en función de la posición del sol y la altura de los edificios. Se continuará descartando aquellas rutas que, aunque seán las más rápidas, superen un umbral de riesgo (muchas zonas sin sombra, sin bancos o fuentes de agua, con muchos desniveles pronunciados...). Finalmente se generará la mejor ruta posible para evitar riesgos.

#### ¿Se ha incluido la configuración del repositorio y se ha enlazado desde el `README`?

Si, por supuesto. La configuración se encuentra arriba. También se ha enlazado con las fotografías del juego de rol.

#### ¿El estudiante tiene todos los datos necesarios para poder resolver el problema, o va a requerir que el usuario los introduzca?

Si, el usuario apenas debe introducir datos, tan solo su ubicación y el lugar al que planea ir. Los datos necesarios están en páginas como el ayuntamiento de Granada o el Instituto Geográfico Nacional.



