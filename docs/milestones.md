# MILESTONES

## [M0] Primera Heurística - Datos

Obtener fichero GeoJSON con los datos referentes a las calles de Atarfe y su número de árboles asociado (utilizando Overpass). Es la primera heurística del sistema, por tanto es imprescindible que en el fichero quede bien reflejada la relación entre las calles y los árboles que hay en cada una de ellas. Deberá, por tanto, contener todas las calles del pueblo y el número de árboles en cada una de ellas (introducir 0 si no se encuentra ningún árbol en una calle). Un árbol se considera dentro de una calle si contiene la etiqueta "avenue" y su distancia con respecto a la calle es menor o igual a 5 metros, o bien si no contiene la etiqueta pero la distancia es menor o igual a 2 metros. El archivo GeoJSON se divide en "features" (elementos, cada calle), estableciendo en cada uno la geometría (al ser una calle, línea con los diferentes puntos que la componen) y propiedades como el ID asignado en OpenStreetMap, el nombre de la calle y el número de árboles.

## [M1] Primera Heurística - Lógica

Con los datos referentes a las calles de Atarfe y el número de árboles por calle, construir una función que construya la heurística de forma que, dada una calle (entrada de la función), se especifique si esta se considera "buena" o "mala" para evitar el sol (salida de la función). La heurística debe tener en cuenta la longitud de la calle, claramente no es lo mismo 3 árboles en una calle de 5 metros que los mismos árboles en una calle de 20 metros. Por tanto, la densidad de árboles por metro de calle es la regla a seguir para la construcción de la heurística. Finalmente, para verificar la validez de la heurística, se comprobarán los resultados con un conjunto de calles ya evaluados manualmente.
