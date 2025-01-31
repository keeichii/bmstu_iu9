string-trim




;; 1. Процедуры для обрезки строк

;; string-trim-left: Удаляет пробелы в начале строки
;; Время: O(n), Память: O(n)
(define (string-trim-left str)
  (let ((s (if (string? str) (string->list str) str)))
    (let recur ((s s))
      (if (and (not (null? s)) (char-whitespace? (car s)))
          (recur (cdr s))
          (if (string? str) (list->string s) s)))))


;; string-trim-right: Удаляет пробелы в конце строки
;; Время: O(n), Память: O(n)
(define (string-trim-right str)
  (let ((s (if (string? str) (string->list str) str)))
    (let loop ((s (reverse s)))
      (let ((trimmed (string-trim-left s)))
        (if (string? str)
            (list->string (reverse trimmed))
            (reverse trimmed))))))

;; string-trim: Удаляет пробелы с обеих сторон строки
;; Время: O(n), Память: O(n)
(define (string-trim str)
  (string-trim-left (string-trim-right str)))


;; 2. Предикаты: Префикс, суффикс и подстрока

;; string-prefix?: Проверяет, является ли строка a префиксом строки b
;; Время: O(min(m, n)), Память: O(1)
(define (string-prefix? a b)
  (let ((a (if (string? a) (string->list a) a))  
        (b (if (string? b) (string->list b) b)))
    (let loop ((a a) (b b))
      (cond ((null? a) #t)
            ((or (null? b) (not (equal? (car a) (car b)))) #f)
            (else (loop (cdr a) (cdr b)))))))

;; string-suffix?: Проверяет, является ли строка a суффиксом строки b
;; Время: O(min(m, n)), Память: O(n)
(define (string-suffix? a b)
  (let ((a (if (string? a) (string->list a) a))
        (b (if (string? b) (string->list b) b)))
    (string-prefix? (reverse a) (reverse b))))

;; string-infix?: Проверяет, является ли строка a подстрокой строки b
;; Время: O(n*m), Память: O(1)
(define (string-infix? a b)
  (let ((a (if (string? a) (string->list a) a))
        (b (if (string? b) (string->list b) b)))
    (cond ((null? a) #t)
          ((null? b) #f)
          ((string-prefix? a b) #t)
          (else (string-infix? a (cdr b))))))


;; 3. Процедура для разделения строки

;; string-split: Разбивает строку на подстроки по разделителю sep
;; Время: O(n), Память: O(n)
(define (string-split str sep)
  (cond
    ((list? str)
     (split-list str sep))
    ((string? str)
     (split-string str sep))
    (else
     (error "Invalid input type"))))

(define (split-string str sep)
  (define (split-helper s acc curr)
    (let ((sep-len (string-length sep)))
      (cond
        ((string=? s "")
         (reverse (cons (list->string (reverse curr)) acc)))
        ((and (>= (string-length s) sep-len) (string=? (substring s 0 sep-len) sep))
         (split-helper (substring s sep-len) (cons (list->string (reverse curr)) acc) '()))
        (else (split-helper (substring s 1) acc (cons (string-ref s 0) curr))))))
  (split-helper str '() '()))

(define (split-list lst sep)
  (define (split-helper lst acc curr)
    (cond
      ((null? lst)
       (reverse (cons (reverse curr) acc)))
      ((equal? (car lst) sep)
       (split-helper (cdr lst) (cons (reverse curr) acc) '()))
      (else
       (split-helper (cdr lst) acc (cons (car lst) curr)))))
  (split-helper lst '() '()))


;; Тест для string-trim-left
(string-trim-left '(#\tab #\space #\a #\b #\c #\space))  ;; -> (#\a #\b #\c #\space)
(string-trim-left '(#\a #\b #\c))  ;; -> (#\a #\b #\c)

;; Тест для string-trim-right
(string-trim-right '(#\a #\b #\c #\space #\tab))  ;; -> (#\a #\b #\c)
(string-trim-right '(#\a #\b #\c))  ;; -> (#\a #\b #\c)

;; Тест для string-trim
(string-trim '(#\tab #\a #\b #\c #\space))  ;; -> (#\a #\b #\c)
(string-trim '(#\a #\b #\c))  ;; -> (#\a #\b #\c)

;; Тест для string-prefix?
(string-prefix? '(#\a #\b) '(#\a #\b #\c #\d))  ;; -> #t
(string-prefix? '(#\x #\y) '(#\a #\b #\c))  ;; -> #f

;; Тест для string-suffix?
(string-suffix? '(#\c #\d) '(#\a #\b #\c #\d))  ;; -> #t
(string-suffix? '(#\x #\y) '(#\a #\b #\c #\d))  ;; -> #f

;; Тест для string-infix?
(string-infix? '(#\b #\c) '(#\a #\b #\c #\d))  ;; -> #t
(string-infix? '(#\x #\y) '(#\a #\b #\c #\d))  ;; -> #f

;; Тест для string-split
(string-split '(#\a #\b #\c #\d) #\c)  ;; -> ((#\a #\b) (#\d))
(string-split '(#\x #\y #\- #\a #\b) #\-)  ;; -> ((#\x #\y) (#\a #\b))

;Антиплагиат: Похожих посылок не найдено
;Комментарий преподавателя:
