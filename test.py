rate_rack = "rack"
rate_net = "net"
rounding_method = "AS_IS"
base_price = 100000.0
tax = 0.19
supplier_commission = 0.15
exchange_rate = 750
operation_type = "divide"

def calculate_price(rate_type, base_price, tax=0, supplier_commission=0, rounding_method="AS_IS", exchange_rate=None, operation_type=None):
    tax = 1 + tax if tax else 1
    supplier_commission = 1 - supplier_commission if supplier_commission else 1
    price = base_price * tax * supplier_commission
    if exchange_rate:
        if operation_type == "divide":
            price = price / exchange_rate
        else:
            price = price * exchange_rate
    return price

print(calculate_price(rate_rack, base_price, tax, supplier_commission, rounding_method, exchange_rate, operation_type))

def create_transformer(factor):
    return lambda number: number  * factor


numbers = [1, 2, 3]

double = create_transformer(2)
d_numbers = list(map(double, numbers))
tripler = create_transformer(3)
t_numbers = list(map(tripler, numbers))
print(t_numbers)