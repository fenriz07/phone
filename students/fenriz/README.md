

4819/5000
# Ejercicio # 8: Normalizador de número de teléfono

[! [estado del ejercicio: publicado] (https://img.shields.io/badge/exercise%20status-released-green.svg?style=for-the-badge)] (https://gophercises.com/exercises /teléfono)

## Detalles del ejercicio

Este ejercicio es bastante sencillo: vamos a escribir un programa que iterará a través de una base de datos y normalizará todos los números de teléfono en la base de datos. Después de normalizar todos los datos, podemos encontrar que hay duplicados, por lo que eliminaremos esos duplicados manteniendo solo una entrada en nuestra base de datos.

Si bien el objetivo principal es crear un programa que normalice la base de datos, también veremos cómo configurar la base de datos y escribir entradas en ella utilizando Go también. Esto se hace intencionalmente para tratar de enseñar cómo usar Go al interactuar con una base de datos SQL.

Hay muchas formas de usar SQL en Go. En lugar de solo elegir uno, voy a tratar de cubrir algunos. Si desea ver cualquier biblioteca adicional cubierta, no dude en comunicarse e intentaré agregarla. Por ahora, aquí están las bibliotecas que pretendo cubrir:

- Escribir SQL sin formato y usar el paquete [database / sql] (https://golang.org/pkg/database/sql/) en la biblioteca estándar
- Usando el muy popular [sqlx] (https://github.com/jmoiron/sqlx) paquete de terceros, que es básicamente una extensión del paquete sql de Go.
- Usando un ORM relativamente minimalista (usaré [gorm] (https://github.com/jinzhu/gorm))

También necesitaremos explorar algunas técnicas básicas de manipulación de cadenas para normalizar nuestros números de teléfono, pero el enfoque principal aquí estará en el SQL, así que intentaré mantener esa parte del código relativamente simple.

Tengo la intención de usar SQLite y Postgres para los videos, pero si desea ver un ejemplo de MySQL (debería ser casi idéntico a los otros dos), avíseme e intentaré agregarlo una vez que el resto de Gophercises están hechos.

En cuanto al ejercicio, comenzaremos creando una base de datos junto con una tabla `phone_numbers`. Dentro de esa tabla queremos agregar las siguientes entradas (sí, sé que hay duplicados):

`` `
1234567890
123 456 7891
(123) 456 7892
(123) 456-7893
123-456-7894
123-456-7890
1234567892
(123)456-7892
`` `

Puede organizar su tabla como desee, y puede agregar los campos adicionales que desee. Mis tablas probablemente variarán según la biblioteca que esté usando, ya que los ORM como GORM a menudo agregan algunos campos automáticamente. También puede crear la tabla manualmente (a través de SQL sin formato) si lo desea, pero intente insertar las entradas con el código Go.

Una vez que haya creado las entradas, nuestro siguiente paso es aprender a iterar sobre las entradas en la base de datos usando el código Go. Con esto deberíamos poder recuperar cada número para poder comenzar a normalizar su contenido.

Una vez que tenga todos los datos en la base de datos, nuestro siguiente paso es normalizar el número de teléfono. Vamos a actualizar todos nuestros números para que coincidan con el formato:

`` `
##########
`` `

Es decir, vamos a eliminar todo el formato y solo almacenar los dígitos. Cuando deseamos mostrar números más tarde, siempre podemos formatearlos, pero por ahora solo necesitamos los dígitos.

En la lista de números proporcionados, la primera entrada, junto con la penúltima entrada, coincide con este formato. Todos los demás no lo hacen y deberán ser reformateados. También hay algunos duplicados que aparecerán una vez que hayamos formateado todos los números, y será necesario eliminarlos de la base de datos, pero no se preocupe por eso por ahora.

Una vez que haya escrito el código que tomará con éxito un número con cualquier formato y devolverá el mismo número en el formato adecuado, utilizaremos una 'ACTUALIZACIÓN' para modificar las entradas en la base de datos. Si el valor que estamos insertando en nuestra base de datos ya existe (es un duplicado), en su lugar, eliminaremos la entrada original.

Cuando termine su programa, las entradas de su base de datos deben verse así (el orden es irrelevante, pero los duplicados deben eliminarse):

`` `
1234567890
1234567891
1234567892
1234567893
1234567894
`` `


## Bono

No hay una bonificación concreta para este ejercicio. Si lo desea, puede explorar otras bibliotecas SQL de terceros y otras bases de datos SQL (MySQL, etc.), pero eso depende de usted.

## Recursos adicionales

He escrito un poco sobre Go y PostgreSQL aquí: https://www.calhoun.io/using-postgresql-with-golang/
Si bien la mayoría de los artículos tratan sobre PostgreSQL, el uso de otras variantes de SQL tiende a ser casi idéntico con algunas excepciones menores (como cuando se conecta a la base de datos o usa características específicas de DB).

No todos los artículos están completos, pero si te atascas, échales un vistazo para obtener ayuda. En particular, si está utilizando Postgres y obtiene un error como `pq: SSL no está habilitado en el servidor` cuando intenta conectarse, le recomiendo consultar este artículo: https://www.calhoun.io/connecting-to-a -postgresql-database-with-gos-database-sql-package /
