linear-equation


(define (linear-equation a11 a12 a21 a22 b1 b2)
  (let* ((det (- (* a11 a22) (* a12 a21)))
         (det-x (- (* b1 a22) (* b2 a12)))
         (det-y (- (* a11 b2) (* a21 b1))))
    (if (not (= det 0))
        (list (/ det-x det) (/ det-y det))
        (list))))

;Антиплагиат: Найдены очень похожие посылки
;Комментарий преподавателя:
