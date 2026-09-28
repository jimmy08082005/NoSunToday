# MILESTONES

## [M0] Primera Heurística - Datos

Existe un fichero GeoJSON con los datos referentes a las calles de Atarfe y su número de árboles asociado. El archivo contiene todas las calles del pueblo (incluyendo calles peatonales) y el número de árboles en cada una de ellas (0 si no hay ningún árbol). Esta información es contrastada manualmente con un conjunto de calles de las cuales se conoce cuántos árboles tienen. Ahora, con estos datos se puede decidir qué calles pueden evitar mejor el sol.

## [M1] Primera Heurística - Lógica

Existe un módulo de código que dice si una calle es "buena" o "mala" para resguardarse frente al sol, en función del número de árboles. Para verificar la validez, se comprueban los resultados con un conjunto de calles ya evaluadas por Manuel, el cual considera una calle "buena" si los árboles dan sombra en la mayor parte de la misma. El módulo se construye sobre el fichero obtenido en el milestone 0.
