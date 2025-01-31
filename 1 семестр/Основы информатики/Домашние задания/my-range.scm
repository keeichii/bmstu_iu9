my-range


(define (my-gcd a b)
  (if (= b 0)
      (abs a)
      (my-gcd b (remainder a b))))

(define (my-lcm a b)
  (/ (* a b) (my-gcd a b)))

(define (prime? n)
  (define (divides? a b)
    (= (remainder b a) 0))
  (define (test-divisor d)
    (cond ((> (* d d) n) #t)
          ((divides? d n) #f)
          (else (test-divisor (+ d 1)))))
  (if (< n 2)
      #f
      (test-divisor 2)))

;Антиплагиат: Найдены похожие посылки
;Комментарий преподавателя:
