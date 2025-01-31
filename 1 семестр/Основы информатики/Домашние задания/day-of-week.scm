day-of-week


(define (day-of-week d m y)
  (let* ((y0 (if (< m 3) (- y 1) y))
         (m0 (if (< m 3) (+ m 12) m))
         (k (remainder y0 100))
         (j (quotient y0 100))
         (h (remainder (+ d (quotient (* 13 (+ m0 1)) 5) k (quotient k 4) (quotient j 4) (- (* 2 j))) 7)))
    (remainder (+ h 6) 7)))

;Антиплагиат: Найдены очень похожие посылки
;Комментарий преподавателя:
