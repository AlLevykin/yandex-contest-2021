# yandex-contest-2021
## Задание № 3

## D. Матрица поворота
Ограничение времени	2 секунды
Ограничение памяти	256Mb
Ввод	стандартный ввод или input.txt
Вывод	стандартный вывод или output.txt
Преобразование Фурье больше не отвечает современным потребностям. Поэтому отдел разработки новых алгоритмов сжатия придумал новое линейное ортогональное преобразование исходных данных, которое улучшает показатель сжатия. К сожалению, последняя строка матрицы этого преобразования утеряна, нужно её восстановить.
Напомним, что линейное преобразование называется ортогональным, если оно сохраняет длину векторов. Кроме того, известно, что исходное преобразование имело определитель равный единице.

Формат ввода
В первой строке входных данных записано число
n
(
2
≤
n
≤
1
6
) — размер матрицы. В следующих
n
−
1
строках записаны по
n
чисел в каждой — первые
N
−
1
строк матрицы преобразования. Элементы матрицы записаны с точностью
1
0
−
1
2
.
Гарантируется, что решение всегда существует.

Формат вывода
Выведите
N
чисел — последнюю строку матрицы. Ответ будет засчитан, если относительная или абсолютная погрешность каждого из чисел не превосходит
1
0
−
6
.
Если линейных ортогональных преобразований, удовлетворяющих условию задачи, несколько, то выведите любое из них.