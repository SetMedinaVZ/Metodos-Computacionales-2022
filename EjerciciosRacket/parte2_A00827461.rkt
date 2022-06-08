#lang racket

;1. insert = funcion que recibe una lista en orden de menor a mayor y un numero el cual va a insertar en su posición correspondiente
(define (insert lst n)
  (cond
    [(empty? lst) (list n)]
    [(<= n (first lst)) (cons n lst)]
    [else (cons (first lst)(insert (rest lst) n))]))

(insert '(1 3 6 7 9 16) 5)

;2. insertion-sort = ordena la lista recibida en orden ascendente
(define (insertion-sort lst) 
  (define (insert lst n)
    (cond
      [(empty? lst) (list n)]
      [(<= n (first lst)) (cons n lst)]
      [else (cons (first lst)(insert (rest lst) n))]))
  (cond
    [(empty? lst) lst]
    [else (insert (insertion-sort(cdr lst)) (first lst))]))

(insertion-sort '(4 3 6 8 3 0 9 1 7))

;5. GCD = retorna el maximo comun divisor de dos numeros
(define (gcd a b)
    (cond
         [(> a b) (gcd b (- a b))]
         [(< a b) (gcd a (- b a))]
         [else a]))
(gcd 6307 1995)

;6. deep-reverse = invierte los elementos de una lista aun con listas anidadas SE CREA LA FUNCION MAP
(define (deep-reverse lst)
  
  (define (map func lst)
    (if (empty? lst)
        '()
        (cons (func (first lst)) (map func (rest lst)))))
  
  (if (list? lst)
      (reverse (map deep-reverse lst))
      lst))

(deep-reverse '(a (b (c (d (e (f (g (h i j)))))))))

;7. insert-everywhere = devuelve todas las combinaciones de una lista con un objeto x
(define (insert-everywhere x lst)
  
  (define (reverseappend lst tail)
  (if (null? lst)
      tail
      (reverseappend (rest lst)
                 (list* (first lst) tail))))
  
  (let f ((tail lst)
           (rhead '())
           (results '()))
    (if (null? tail)
        (reverse (list* (reverseappend rhead (list* x tail)) results))
        (f (rest tail)
            (list* (first tail) rhead)
            (list* (reverseappend rhead (list* x tail)) results)))))

(insert-everywhere 'x '(1 2 3 4 5 6 7 8 9 10))

;18. Integral por metodo de simpson
(define (integral a b n f)
  (define (inc x) (+ x 1))
  (define h (/ (- b a) n))
  (define (yk k) (f (+ a (* h k))))
  (define (sum term a next b)
  (if (> a b)
      0
      (+ (term a)
         (sum term (next a) next b))))
  (define (simpson-term k)
    (* (cond
         [(or (= k 0) (= k n)) 1]
         [(odd? k) 4]
         [else 2])
       (yk k)))
  (* (/ h 3) (sum simpson-term 0 inc n)))

(integral 0 1 10 (lambda (x) (* x x x)))