list->set


;; Определение функции remove: удаляет все вхождения элемента x из списка xs
(define (remove x xs)
  (cond
    ((null? xs) '())
    ((equal? x (car xs)) (remove x (cdr xs)))
    (else (cons (car xs) (remove x (cdr xs))))))


;; list->set: Преобразует список xs в множество, удаляя дубликаты
;; Время: O(n^2), Память: O(n)
(define (list->set xs)
  (if (null? xs)
      '()
      (cons (car xs) (list->set (remove (car xs) (cdr xs))))))

;; set?: Проверяет, является ли список xs множеством (содержит ли он уникальные элементы)
;; Время: O(n^2), Память: O(n)
(define (set? xs)
  (cond
    ((null? xs) #t)
    ((member (car xs) (cdr xs)) #f)
    (else (set? (cdr xs)))))

;; union: Возвращает объединение множеств xs и ys
;; Время: O(n^2), Память: O(n)
(define (union xs ys)
  (if (null? xs)
      ys
      (cons (car xs) (union (remove (car xs) ys) (cdr xs)))))

;; intersection: Возвращает пересечение множеств xs и ys
;; Время: O(n^2), Память: O(n)
(define (intersection xs ys)
  (cond
    ((null? xs) '())
    ((member (car xs) ys) (cons (car xs) (intersection (cdr xs) ys)))
    (else (intersection (cdr xs) ys))))

;; difference: Возвращает разность множеств xs и ys
;; Время: O(n^2), Память: O(n)
(define (difference xs ys)
  (cond
    ((null? xs) '())
    ((member (car xs) ys) (difference (cdr xs) ys))
    (else (cons (car xs) (difference (cdr xs) ys)))))

;; symmetric-difference: Возвращает симметричную разность множеств xs и ys
;; Время: O(n^2), Память: O(n)
(define (symmetric-difference xs ys)
  (union (difference xs ys) (difference ys xs)))

;; set-eq?: Проверяет, равны ли множества xs и ys
;; Время: O(n^2), Память: O(n)
(define (set-eq? xs ys)
  (and (null? (difference xs ys)) (null? (difference ys xs))))


;; Тесты для list->set
(list->set '(1 1 2 3)) ;; -> (1 2 3)
(list->set '(a b c a b)) ;; -> (a b c)
(list->set '(5 4 4 4 3 2 1)) ;; -> (5 4 3 2 1)

;; Тесты для set?
(set? '(1 2 3)) ;; -> #t
(set? '(1 2 2 3)) ;; -> #f
(set? '()) ;; -> #t

;; Тесты для union
(union '(1 2 3) '(3 4 5)) ;; -> (1 2 3 4 5)
(union '(a b) '(b c d)) ;; -> (a b c d)
(union '(1 2) '(3 4)) ;; -> (1 2 3 4)

;; Тесты для intersection
(intersection '(1 2 3) '(2 3 4)) ;; -> (2 3)
(intersection '(a b c) '(b c d)) ;; -> (b c)
(intersection '(1 2 3) '(4 5 6)) ;; -> ()

;; Тесты для difference
(difference '(1 2 3) '(2 3 4)) ;; -> (1)
(difference '(a b c) '(b c d)) ;; -> (a)
(difference '(1 2 3) '(1 2 3)) ;; -> ()

;; Тесты для symmetric-difference
(symmetric-difference '(1 2 3) '(2 3 4)) ;; -> (1 4)
(symmetric-difference '(a b c) '(b c d)) ;; -> (a d)
(symmetric-difference '(1 2 3) '(1 2 3)) ;; -> ()

;; Тесты для set-eq?
(set-eq? '(1 2 3) '(3 2 1)) ;; -> #t
(set-eq? '(a b c) '(a b d)) ;; -> #f
(set-eq? '(1 2) '(1 2 3)) ;; -> #f

;Антиплагиат: Найдены похожие посылки
;Комментарий преподавателя: Достаточно было написать только оценки времени. Процедура set? не создаёт cons-ячеек, рекурсия хвостовая (внутри встроенной member, скорее всего, тоже, либо она вообще написана через цикл на Си), так что её потребление памяти O(1).
