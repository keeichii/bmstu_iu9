multi-vector


;; Реализация fold-left
(define (fold-left f init lst)
  (let loop ((acc init)
             (lst lst))
    (if (null? lst)
        acc
        (loop (f acc (car lst)) (cdr lst)))))

;; make-multi-vector: создает многомерный вектор с заданным размером и заполняет значением fill
(define (make-multi-vector sizes . fill)
  (let* ((total-size (apply * sizes))
         (data (make-vector total-size (if (null? fill) #f (car fill)))))
    (vector sizes data)))

;; multi-vector?: проверяет, является ли структура многомерным вектором
(define (multi-vector? m)
  (and (vector? m)
       (= (vector-length m) 2)
       (vector? (vector-ref m 1))
       (list? (vector-ref m 0))))

;; multi-vector-ref: получает значение из многомерного вектора по индексам
(define (multi-vector-ref m indices)
  (let* ((sizes (vector-ref m 0))
         (data (vector-ref m 1))
         (flat-index (fold-left
                      (lambda (acc pair)
                        (+ (* acc (car pair)) (cdr pair)))
                      0
                      (map cons sizes indices))))
    (vector-ref data flat-index)))

;; multi-vector-set!: изменяет значение элемента многомерного вектора по индексам
(define (multi-vector-set! m indices value)
  (let* ((sizes (vector-ref m 0))
         (data (vector-ref m 1))
         (flat-index (fold-left
                      (lambda (acc pair)
                        (+ (* acc (car pair)) (cdr pair)))
                      0
                      (map cons sizes indices))))
    (vector-set! data flat-index value)))

;; Тесты для make-multi-vector
(define m (make-multi-vector '(3 4)))
(multi-vector? m) ;; -> #t
(define m2 (make-multi-vector '(2 2 2) 0))
(multi-vector-ref m2 '(1 1 1)) ;; -> 0

;; Тесты для make-multi-vector с fill
(define m3 (make-multi-vector '(2 3) -1))
(multi-vector-ref m3 '(1 2)) ;; -> -1
(define m4 (make-multi-vector '(2 2 2) 5))
(multi-vector-ref m4 '(0 0 1)) ;; -> 5

;; Тесты для multi-vector?
(multi-vector? (make-multi-vector '(2 2))) ;; -> #t
(multi-vector? '(1 2 3)) ;; -> #f

;; Тесты для multi-vector-ref
(define m5 (make-multi-vector '(2 2 2) 0))
(multi-vector-ref m5 '(1 0 1)) ;; -> 0
(define m6 (make-multi-vector '(1 1 1) 'test))
(multi-vector-ref m6 '(0 0 0)) ;; -> 'test

;; Тесты для multi-vector-set!
(define m7 (make-multi-vector '(2 2 2) 0))
(multi-vector-set! m7 '(1 1 1) 'new-value)
(multi-vector-ref m7 '(1 1 1)) ;; -> 'new-value
(define m8 (make-multi-vector '(3 3 3) 'old-value))
(multi-vector-set! m8 '(2 2 2) 'changed)
(multi-vector-ref m8 '(2 2 2)) ;; -> 'changed

;Антиплагиат: Найдены похожие посылки
;Комментарий преподавателя:
