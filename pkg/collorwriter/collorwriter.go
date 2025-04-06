package collorwriter

/*
	В данном коде используется пакет fmt для форматирования строк и пакет bytes для работы с байтовыми данными.
   	ColorWriter - это структура, которая содержит поле w типа io.Writer, которое будет использоваться для записи данных.
	Write - это метод, который принимает массив байтов p и возвращает количество записанных байтов n и ошибку err.
	Строка разбивается на пробелы и первые 3 пробела красятся в цвета синего и желтого соответственно.
*/
import (
	"bytes"
	"fmt"
	"io"
)

type ColorWriter struct {
	W io.Writer
}

func (c *ColorWriter) Write(p []byte) (n int, err error) {
	parts := bytes.SplitN(p, []byte{' '}, 3)
	if len(parts) < 3 {
		return c.W.Write(p)
	}
	colored := fmt.Sprintf(
		"\033[34m%s %s\033[0m \033[33m%s\033[0m",
		parts[0], parts[1], parts[2],
	)
	return c.W.Write([]byte(colored))
}
