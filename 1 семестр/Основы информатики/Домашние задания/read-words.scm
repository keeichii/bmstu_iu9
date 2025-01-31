read-words


(define (read-words)
  (define (read-word port)
    (let loop ((chars '()))
      (let ((ch (peek-char port)))
        (cond
          ((or (eof-object? ch) (char-whitespace? ch))
           (if (null? chars)
               (begin
                 (read-char port)  ; пропускаем текущий пробельный символ
                 (if (eof-object? ch) '() (loop chars)))
               (reverse chars)))
          (else
           (read-char port)
           (loop (cons ch chars)))))))

  (define (read-words-loop port words)
    (let ((word (read-word port)))
      (if (null? word)
          (if (eof-object? (peek-char port)) (reverse words)
              (read-words-loop port words))
          (read-words-loop port (cons (list->string word) words)))))

  (read-words-loop (current-input-port) '()))


;Антиплагиат: Найдены похожие посылки
;Комментарий преподавателя:
