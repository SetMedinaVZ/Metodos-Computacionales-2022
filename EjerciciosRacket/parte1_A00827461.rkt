#lang racket
;7. Pow Function
(define (pow base n)
  (cond ((or (= base 1) (= n 0)) 1)
        (else (* base (pow base (- n 1))))))

;8. Fib member function -> empieza desde 0
(define (fiboMember n)
  (if (<= n 2)
      (+ 1 0)
      (+ (fiboMember(- n 1))(fiboMember(- n 2)))))

;10 Positives Function -> regresa una lista con los numeros positivos
(define (positives lst)
  (define (iter lst accumulator)
    (cond
      [(empty? lst) (reverse accumulator)]
      [(>= (first lst) 0) (iter (rest lst) (cons (first lst) accumulator))]
      [else (iter (rest lst) accumulator)]))
  (iter lst empty))
      
;11. add-list Function -> regresa la suma de los numeros contenidos en la lista
(define (add-list lst)
  (if (empty? lst) 0
      (+ (first lst) (add-list (rest lst)))))

;13 list-ofsymbols? -> regresa t/f si los elementos en lst son simbolos
(define (list-of-symbols? lst)
  (define (check lst b)
    (cond
      [(empty? lst) b]
      [else
       (cond
         [(symbol? (first lst)) (check (rest lst) #t )]
         [else (check empty #f)])]))
  (check lst empty))

;15. dot-product -> devuelve el producto punto a a por b
(define (dot-product a b)
  (cond ((null? a) 0)
        (else
         (+ (* (car a) (car b))
            (dot-product (cdr a) (cdr b))))))
 

;16 average
(define (average xs (return /))
  (if (empty? xs)
      (return 0 0)
      (average (cdr xs)
               (lambda (sum len)
                 (return (+ sum (car xs))
                         (+ len 1))))))

;19 expand1
(define (expand lst)
  (define lth (length lst))
  (define (iter lth lst a x)
    (cond
      [(empty? lst) (reverse a)]
      [(iter lth (rest lst) (repeat x (first lst) a) (+ x 1))]))

  (define (repeat x first a)
    (cond
      [(= x 0) a]
      [else (repeat (- x 1) first (cons first a))]))

  (iter lth lst empty 1))



