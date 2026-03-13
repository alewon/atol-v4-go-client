# API АТОЛ Онлайн v4: методы и поля

Источник истины: официальный PDF `API сервиса АТОЛ Онлайн_v4`.

Дополнительно 13.03.2026 выполнена живая проверка на тестовом стенде:

- `https://testonline.atol.ru/possystem/v4`
- `group_code = v4-online-atol-ru_4179`
- без проверки `callback_url`

Рабочие правила для этого файла:

- перечислены только методы API и webhook `callback_url`, которые описаны в PDF;
- ссылки `https://online.atol.ru/possystem/v4/schema/sell` и `https://online.atol.ru/possystem/v4/schema/correction` использованы только как источник состава полей и не считаются отдельными методами;
- повторяющиеся структуры намеренно дублируются в каждой секции;
- если поведение тестового стенда расходится с PDF, это отмечено прямо в описании поля или в разделе с примечаниями по методу;
- если обязательность поля нельзя было надёжно подтвердить из доступных локально данных, в колонке `required` указано `не указано явно`.

Типы в колонке `type`:

- `string`, `int`, `float`, `bool` для скаляров;
- `object` для вложенного JSON-объекта;
- `array<object>` и `array<string>` для массивов;
- `null` означает, что поле допускает `null`.

## 1. `POST /possystem/v4/getToken`

Полный URL:

```text
https://online.atol.ru/possystem/v4/getToken
```

### Параметры URL

Нет path/query параметров.

### Тело запроса

| field | type | required | description |
|---|---|---|---|
| `login` | `string` | да | Логин пользователя API. |
| `pass` | `string` | да | Пароль пользователя API. |
| `source` | `string` | не указано явно | Идентификатор источника запроса, если используется интеграцией. |

### Успешный ответ

| field | type | required | description |
|---|---|---|---|
| `token` | `string` | да | Токен доступа для последующих вызовов API. |
| `error` | `object \| null` | да | Ошибка. В успешном ответе обычно `null`. |
| `timestamp` | `string` | да | Время формирования ответа. |

### Ответ с ошибкой

В доступных локально данных отдельная error-схема для `POST /getToken` не выделена. Подтверждён только ответ с полями:

| field | type | required | description |
|---|---|---|---|
| `token` | `string` | не указано явно | Токен. При ошибке может отсутствовать или быть пустым. |
| `error.error_id` | `string` | не указано явно | Идентификатор ошибки. |
| `error.code` | `int` | не указано явно | Код ошибки. |
| `error.text` | `string` | не указано явно | Текст ошибки. |
| `error.type` | `string` | не указано явно | Тип ошибки. |
| `timestamp` | `string` | не указано явно | Время формирования ответа. |

## 2. `GET /possystem/v4/getToken?login=<login>&pass=<pass>`

Полный URL:

```text
https://online.atol.ru/possystem/v4/getToken?login=<login>&pass=<pass>
```

### Параметры URL

| field | type | required | description |
|---|---|---|---|
| `login` | `string` | да | Логин пользователя API. |
| `pass` | `string` | да | Пароль пользователя API. |
| `source` | `string` | не указано явно | Идентификатор источника запроса, если используется интеграцией. |

### Тело запроса

Отсутствует.

### Успешный ответ

| field | type | required | description |
|---|---|---|---|
| `token` | `string` | да | Токен доступа для последующих вызовов API. |
| `error` | `object \| null` | да | Ошибка. В успешном ответе обычно `null`. |
| `timestamp` | `string` | да | Время формирования ответа. |

### Ответ с ошибкой

В доступных локально данных отдельная error-схема для `GET /getToken` не выделена. Подтверждён только ответ с полями:

| field | type | required | description |
|---|---|---|---|
| `token` | `string` | не указано явно | Токен. При ошибке может отсутствовать или быть пустым. |
| `error.error_id` | `string` | не указано явно | Идентификатор ошибки. |
| `error.code` | `int` | не указано явно | Код ошибки. |
| `error.text` | `string` | не указано явно | Текст ошибки. |
| `error.type` | `string` | не указано явно | Тип ошибки. |
| `timestamp` | `string` | не указано явно | Время формирования ответа. |

## 3. `POST /possystem/v4/{group_code}/{operation}`

Полный URL:

```text
https://online.atol.ru/possystem/v4/{group_code}/{operation}
```

Дополнительно в PDF показан вариант:

```text
https://online.atol.ru/possystem/v4/{group_code}/{operation}?token=<token>
```

### Общие параметры URL

| field | type | required | description |
|---|---|---|---|
| `group_code` | `string` | да | Код группы ККТ. |
| `operation` | `string` | да | Операция регистрации документа. |
| `token` | `string` | не указано явно | Токен в query string, если используется такой вариант авторизации. |

Допустимые `operation`:

- `sell`
- `sell_refund`
- `buy`
- `buy_refund`
- `sell_correction`
- `buy_correction`

### 3.1 `POST /possystem/v4/{group_code}/sell`

#### Тело запроса

| field | type | required | description |
|---|---|---|---|
| `timestamp` | `string` | да | Время формирования документа. |
| `source_id` | `int` | не указано явно | Числовой идентификатор источника. |
| `external_id` | `string` | да | Внешний идентификатор документа у интегратора. |
| `service` | `object` | не указано явно | Служебные параметры обработки. |
| `service.callback_url` | `string` | не указано явно | URL webhook, на который АТОЛ отправит результат обработки. |
| `receipt` | `object` | да | Тело чека прихода. |
| `receipt.client` | `object` | да | Данные покупателя. |
| `receipt.client.name` | `string` | не указано явно | Имя или наименование покупателя. |
| `receipt.client.inn` | `string` | не указано явно | ИНН покупателя. |
| `receipt.client.email` | `string` | условно | Email покупателя. В объекте `receipt.client` обязательно одно из полей: `email` или `phone`. |
| `receipt.client.phone` | `string` | условно | Телефон покупателя. В объекте `receipt.client` обязательно одно из полей: `email` или `phone`. |
| `receipt.company` | `object` | да | Данные организации-продавца. |
| `receipt.company.email` | `string` | не указано явно | Email организации. |
| `receipt.company.sno` | `string` | не указано явно | Система налогообложения. |
| `receipt.company.inn` | `string` | да | ИНН организации. |
| `receipt.company.payment_address` | `string` | да | Адрес расчётов или сайт. |
| `receipt.company.location` | `string` | не указано явно | Место расчёта. |
| `receipt.agent_info` | `object` | не указано явно | Сведения об агенте на уровне чека. |
| `receipt.agent_info.type` | `string` | не указано явно | Тип агента. |
| `receipt.agent_info.paying_agent` | `object` | не указано явно | Данные платёжного агента. |
| `receipt.agent_info.paying_agent.operation` | `string` | не указано явно | Операция агента. |
| `receipt.agent_info.paying_agent.phones` | `array<string>` | не указано явно | Телефоны платёжного агента. |
| `receipt.agent_info.receive_payments_operator` | `object` | не указано явно | Данные оператора по приёму платежей. |
| `receipt.agent_info.receive_payments_operator.phones` | `array<string>` | не указано явно | Телефоны оператора по приёму платежей. |
| `receipt.agent_info.money_transfer_operator` | `object` | не указано явно | Данные оператора перевода денежных средств. |
| `receipt.agent_info.money_transfer_operator.phones` | `array<string>` | не указано явно | Телефоны оператора перевода денежных средств. |
| `receipt.agent_info.money_transfer_operator.name` | `string` | не указано явно | Наименование оператора перевода денежных средств. |
| `receipt.agent_info.money_transfer_operator.address` | `string` | не указано явно | Адрес оператора перевода денежных средств. |
| `receipt.agent_info.money_transfer_operator.inn` | `string` | не указано явно | ИНН оператора перевода денежных средств. |
| `receipt.supplier_info` | `object` | не указано явно | Сведения о поставщике на уровне чека. |
| `receipt.supplier_info.phones` | `array<string>` | не указано явно | Телефоны поставщика. |
| `receipt.supplier_info.name` | `string` | не указано явно | Наименование поставщика. |
| `receipt.supplier_info.inn` | `string` | не указано явно | ИНН поставщика. |
| `receipt.items` | `array<object>` | да | Позиции чека. |
| `receipt.items[].name` | `string` | да | Наименование позиции. |
| `receipt.items[].price` | `float` | да | Цена за единицу. |
| `receipt.items[].quantity` | `float` | да | Количество. |
| `receipt.items[].sum` | `float \| null` | да | Сумма по позиции. |
| `receipt.items[].measurement_unit` | `string` | не указано явно | Единица измерения. |
| `receipt.items[].payment_method` | `string` | не указано явно | Признак способа расчёта. |
| `receipt.items[].payment_object` | `string` | не указано явно | Признак предмета расчёта. |
| `receipt.items[].nomenclature_code` | `string` | не указано явно | Код товара или номенклатуры. |
| `receipt.items[].vat` | `object` | не указано явно | НДС по позиции. |
| `receipt.items[].vat.type` | `string` | да | Тип ставки НДС. |
| `receipt.items[].vat.sum` | `float \| null` | да | Сумма НДС. |
| `receipt.items[].agent_info` | `object` | не указано явно | Сведения об агенте на уровне позиции. |
| `receipt.items[].agent_info.type` | `string` | не указано явно | Тип агента. |
| `receipt.items[].agent_info.paying_agent` | `object` | не указано явно | Данные платёжного агента. |
| `receipt.items[].agent_info.paying_agent.operation` | `string` | не указано явно | Операция агента. |
| `receipt.items[].agent_info.paying_agent.phones` | `array<string>` | не указано явно | Телефоны платёжного агента. |
| `receipt.items[].agent_info.receive_payments_operator` | `object` | не указано явно | Данные оператора по приёму платежей. |
| `receipt.items[].agent_info.receive_payments_operator.phones` | `array<string>` | не указано явно | Телефоны оператора по приёму платежей. |
| `receipt.items[].agent_info.money_transfer_operator` | `object` | не указано явно | Данные оператора перевода денежных средств. |
| `receipt.items[].agent_info.money_transfer_operator.phones` | `array<string>` | не указано явно | Телефоны оператора перевода денежных средств. |
| `receipt.items[].agent_info.money_transfer_operator.name` | `string` | не указано явно | Наименование оператора перевода денежных средств. |
| `receipt.items[].agent_info.money_transfer_operator.address` | `string` | не указано явно | Адрес оператора перевода денежных средств. |
| `receipt.items[].agent_info.money_transfer_operator.inn` | `string` | не указано явно | ИНН оператора перевода денежных средств. |
| `receipt.items[].supplier_info` | `object` | не указано явно | Сведения о поставщике на уровне позиции. |
| `receipt.items[].supplier_info.phones` | `array<string>` | не указано явно | Телефоны поставщика. |
| `receipt.items[].supplier_info.name` | `string` | не указано явно | Наименование поставщика. |
| `receipt.items[].supplier_info.inn` | `string` | не указано явно | ИНН поставщика. |
| `receipt.items[].user_data` | `string` | не указано явно | Дополнительные пользовательские данные по позиции. |
| `receipt.items[].excise` | `float` | не указано явно | Сумма акциза. |
| `receipt.items[].country_code` | `string` | не указано явно | Цифровой код страны происхождения. |
| `receipt.items[].declaration_number` | `string` | не указано явно | Номер таможенной декларации. |
| `receipt.payments` | `array<object>` | да | Сведения об оплате. |
| `receipt.payments[].type` | `int` | да | Тип оплаты. |
| `receipt.payments[].sum` | `float \| null` | да | Сумма оплаты. |
| `receipt.vats` | `array<object>` | не указано явно | Итоговые НДС по чеку. |
| `receipt.vats[].type` | `string` | да | Тип ставки НДС. |
| `receipt.vats[].sum` | `float \| null` | да | Сумма НДС. |
| `receipt.total` | `float` | да | Общая сумма чека. |
| `receipt.additional_check_props` | `string` | не указано явно | Дополнительные реквизиты чека. |
| `receipt.cashier` | `string` | не указано явно | Кассир. |
| `receipt.additional_user_props` | `object` | не указано явно | Дополнительные пользовательские реквизиты. |
| `receipt.additional_user_props.name` | `string` | да | Имя дополнительного реквизита. |
| `receipt.additional_user_props.value` | `string` | да | Значение дополнительного реквизита. |
| `receipt.device_number` | `string` | не указано явно | Номер устройства. |
| `receipt.internet` | `bool` | не указано явно | Признак интернет-расчёта. |
| `receipt.cashless_payments` | `array<object>` | не указано явно | Дополнительные сведения по безналичным оплатам. |
| `receipt.cashless_payments[].sum` | `float` | да | Сумма безналичной оплаты. |
| `receipt.cashless_payments[].method` | `int` | да | Метод безналичной оплаты. |
| `receipt.cashless_payments[].id` | `string` | да | Идентификатор безналичной оплаты. |
| `receipt.cashless_payments[].additional_info` | `string` | не указано явно | Дополнительная информация по безналичной оплате. |

#### Успешный ответ

| field | type | required | description |
|---|---|---|---|
| `uuid` | `string` | да | Идентификатор зарегистрированного документа в системе АТОЛ. |
| `timestamp` | `string` | да | Время формирования ответа. |
| `error` | `object \| null` | да | Ошибка. В успешном ответе обычно `null`. |
| `status` | `string` | не указано явно | Статус обработки документа, например `wait`. |

#### Ответ с ошибкой

| field | type | required | description |
|---|---|---|---|
| `uuid` | `string` | не указано явно | Идентификатор документа, если уже был присвоен. |
| `error` | `object` | да | Объект ошибки. |
| `error.error_id` | `string` | да | Идентификатор ошибки. |
| `error.code` | `int` | да | Код ошибки. |
| `error.text` | `string` | да | Текст ошибки. |
| `error.type` | `string` | не указано явно | Тип ошибки. |
| `timestamp` | `string` | да | Время формирования ответа. |
| `status` | `string` | не указано явно | Статус обработки, например `fail`. |
| `group_code` | `string` | не указано явно | Код группы ККТ. |
| `daemon_code` | `string` | не указано явно | Код демона. |
| `device_code` | `string` | не указано явно | Код устройства. |
| `callback_url` | `string` | не указано явно | Callback URL, если был передан. |
| `payload` | `null` | не указано явно | Полезная нагрузка; в ошибочном ответе в тестах равна `null`. |

### 3.2 `POST /possystem/v4/{group_code}/sell_refund`

#### Тело запроса

| field | type | required | description |
|---|---|---|---|
| `timestamp` | `string` | да | Время формирования документа. |
| `source_id` | `int` | не указано явно | Числовой идентификатор источника. |
| `external_id` | `string` | да | Внешний идентификатор документа у интегратора. |
| `service` | `object` | не указано явно | Служебные параметры обработки. |
| `service.callback_url` | `string` | не указано явно | URL webhook, на который АТОЛ отправит результат обработки. |
| `receipt` | `object` | да | Тело чека возврата прихода. |
| `receipt.client` | `object` | да | Данные покупателя. |
| `receipt.client.name` | `string` | не указано явно | Имя или наименование покупателя. |
| `receipt.client.inn` | `string` | не указано явно | ИНН покупателя. |
| `receipt.client.email` | `string` | условно | Email покупателя. В объекте `receipt.client` обязательно одно из полей: `email` или `phone`. |
| `receipt.client.phone` | `string` | условно | Телефон покупателя. В объекте `receipt.client` обязательно одно из полей: `email` или `phone`. |
| `receipt.company` | `object` | да | Данные организации-продавца. |
| `receipt.company.email` | `string` | не указано явно | Email организации. |
| `receipt.company.sno` | `string` | не указано явно | Система налогообложения. |
| `receipt.company.inn` | `string` | да | ИНН организации. |
| `receipt.company.payment_address` | `string` | да | Адрес расчётов или сайт. |
| `receipt.company.location` | `string` | не указано явно | Место расчёта. |
| `receipt.agent_info` | `object` | не указано явно | Сведения об агенте на уровне чека. |
| `receipt.agent_info.type` | `string` | не указано явно | Тип агента. |
| `receipt.agent_info.paying_agent` | `object` | не указано явно | Данные платёжного агента. |
| `receipt.agent_info.paying_agent.operation` | `string` | не указано явно | Операция агента. |
| `receipt.agent_info.paying_agent.phones` | `array<string>` | не указано явно | Телефоны платёжного агента. |
| `receipt.agent_info.receive_payments_operator` | `object` | не указано явно | Данные оператора по приёму платежей. |
| `receipt.agent_info.receive_payments_operator.phones` | `array<string>` | не указано явно | Телефоны оператора по приёму платежей. |
| `receipt.agent_info.money_transfer_operator` | `object` | не указано явно | Данные оператора перевода денежных средств. |
| `receipt.agent_info.money_transfer_operator.phones` | `array<string>` | не указано явно | Телефоны оператора перевода денежных средств. |
| `receipt.agent_info.money_transfer_operator.name` | `string` | не указано явно | Наименование оператора перевода денежных средств. |
| `receipt.agent_info.money_transfer_operator.address` | `string` | не указано явно | Адрес оператора перевода денежных средств. |
| `receipt.agent_info.money_transfer_operator.inn` | `string` | не указано явно | ИНН оператора перевода денежных средств. |
| `receipt.supplier_info` | `object` | не указано явно | Сведения о поставщике на уровне чека. |
| `receipt.supplier_info.phones` | `array<string>` | не указано явно | Телефоны поставщика. |
| `receipt.supplier_info.name` | `string` | не указано явно | Наименование поставщика. |
| `receipt.supplier_info.inn` | `string` | не указано явно | ИНН поставщика. |
| `receipt.items` | `array<object>` | да | Позиции чека. |
| `receipt.items[].name` | `string` | да | Наименование позиции. |
| `receipt.items[].price` | `float` | да | Цена за единицу. |
| `receipt.items[].quantity` | `float` | да | Количество. |
| `receipt.items[].sum` | `float \| null` | да | Сумма по позиции. |
| `receipt.items[].measurement_unit` | `string` | не указано явно | Единица измерения. |
| `receipt.items[].payment_method` | `string` | не указано явно | Признак способа расчёта. |
| `receipt.items[].payment_object` | `string` | не указано явно | Признак предмета расчёта. |
| `receipt.items[].nomenclature_code` | `string` | не указано явно | Код товара или номенклатуры. |
| `receipt.items[].vat` | `object` | не указано явно | НДС по позиции. |
| `receipt.items[].vat.type` | `string` | да | Тип ставки НДС. |
| `receipt.items[].vat.sum` | `float \| null` | да | Сумма НДС. |
| `receipt.items[].agent_info` | `object` | не указано явно | Сведения об агенте на уровне позиции. |
| `receipt.items[].agent_info.type` | `string` | не указано явно | Тип агента. |
| `receipt.items[].agent_info.paying_agent` | `object` | не указано явно | Данные платёжного агента. |
| `receipt.items[].agent_info.paying_agent.operation` | `string` | не указано явно | Операция агента. |
| `receipt.items[].agent_info.paying_agent.phones` | `array<string>` | не указано явно | Телефоны платёжного агента. |
| `receipt.items[].agent_info.receive_payments_operator` | `object` | не указано явно | Данные оператора по приёму платежей. |
| `receipt.items[].agent_info.receive_payments_operator.phones` | `array<string>` | не указано явно | Телефоны оператора по приёму платежей. |
| `receipt.items[].agent_info.money_transfer_operator` | `object` | не указано явно | Данные оператора перевода денежных средств. |
| `receipt.items[].agent_info.money_transfer_operator.phones` | `array<string>` | не указано явно | Телефоны оператора перевода денежных средств. |
| `receipt.items[].agent_info.money_transfer_operator.name` | `string` | не указано явно | Наименование оператора перевода денежных средств. |
| `receipt.items[].agent_info.money_transfer_operator.address` | `string` | не указано явно | Адрес оператора перевода денежных средств. |
| `receipt.items[].agent_info.money_transfer_operator.inn` | `string` | не указано явно | ИНН оператора перевода денежных средств. |
| `receipt.items[].supplier_info` | `object` | не указано явно | Сведения о поставщике на уровне позиции. |
| `receipt.items[].supplier_info.phones` | `array<string>` | не указано явно | Телефоны поставщика. |
| `receipt.items[].supplier_info.name` | `string` | не указано явно | Наименование поставщика. |
| `receipt.items[].supplier_info.inn` | `string` | не указано явно | ИНН поставщика. |
| `receipt.items[].user_data` | `string` | не указано явно | Дополнительные пользовательские данные по позиции. |
| `receipt.items[].excise` | `float` | не указано явно | Сумма акциза. |
| `receipt.items[].country_code` | `string` | не указано явно | Цифровой код страны происхождения. |
| `receipt.items[].declaration_number` | `string` | не указано явно | Номер таможенной декларации. |
| `receipt.payments` | `array<object>` | да | Сведения об оплате. |
| `receipt.payments[].type` | `int` | да | Тип оплаты. |
| `receipt.payments[].sum` | `float \| null` | да | Сумма оплаты. |
| `receipt.vats` | `array<object>` | не указано явно | Итоговые НДС по чеку. |
| `receipt.vats[].type` | `string` | да | Тип ставки НДС. |
| `receipt.vats[].sum` | `float \| null` | да | Сумма НДС. |
| `receipt.total` | `float` | да | Общая сумма чека. |
| `receipt.additional_check_props` | `string` | не указано явно | Дополнительные реквизиты чека. |
| `receipt.cashier` | `string` | не указано явно | Кассир. |
| `receipt.additional_user_props` | `object` | не указано явно | Дополнительные пользовательские реквизиты. |
| `receipt.additional_user_props.name` | `string` | да | Имя дополнительного реквизита. |
| `receipt.additional_user_props.value` | `string` | да | Значение дополнительного реквизита. |
| `receipt.device_number` | `string` | не указано явно | Номер устройства. |
| `receipt.internet` | `bool` | не указано явно | Признак интернет-расчёта. |
| `receipt.cashless_payments` | `array<object>` | не указано явно | Дополнительные сведения по безналичным оплатам. |
| `receipt.cashless_payments[].sum` | `float` | да | Сумма безналичной оплаты. |
| `receipt.cashless_payments[].method` | `int` | да | Метод безналичной оплаты. |
| `receipt.cashless_payments[].id` | `string` | да | Идентификатор безналичной оплаты. |
| `receipt.cashless_payments[].additional_info` | `string` | не указано явно | Дополнительная информация по безналичной оплате. |

#### Успешный ответ

| field | type | required | description |
|---|---|---|---|
| `uuid` | `string` | да | Идентификатор зарегистрированного документа в системе АТОЛ. |
| `timestamp` | `string` | да | Время формирования ответа. |
| `error` | `object \| null` | да | Ошибка. В успешном ответе обычно `null`. |
| `status` | `string` | не указано явно | Статус обработки документа, например `wait`. |

#### Ответ с ошибкой

| field | type | required | description |
|---|---|---|---|
| `uuid` | `string` | не указано явно | Идентификатор документа, если уже был присвоен. |
| `error` | `object` | да | Объект ошибки. |
| `error.error_id` | `string` | да | Идентификатор ошибки. |
| `error.code` | `int` | да | Код ошибки. |
| `error.text` | `string` | да | Текст ошибки. |
| `error.type` | `string` | не указано явно | Тип ошибки. |
| `timestamp` | `string` | да | Время формирования ответа. |
| `status` | `string` | не указано явно | Статус обработки, например `fail`. |
| `group_code` | `string` | не указано явно | Код группы ККТ. |
| `daemon_code` | `string` | не указано явно | Код демона. |
| `device_code` | `string` | не указано явно | Код устройства. |
| `callback_url` | `string` | не указано явно | Callback URL, если был передан. |
| `payload` | `null` | не указано явно | Полезная нагрузка; в ошибочном ответе в тестах равна `null`. |

### 3.3 `POST /possystem/v4/{group_code}/buy`

#### Тело запроса

Структура тела полностью совпадает с `sell` и дублируется здесь намеренно.

| field | type | required | description |
|---|---|---|---|
| `timestamp` | `string` | да | Время формирования документа. |
| `source_id` | `int` | не указано явно | Числовой идентификатор источника. |
| `external_id` | `string` | да | Внешний идентификатор документа у интегратора. |
| `service` | `object` | не указано явно | Служебные параметры обработки. |
| `service.callback_url` | `string` | не указано явно | URL webhook, на который АТОЛ отправит результат обработки. |
| `receipt` | `object` | да | Тело чека расхода. |
| `receipt.client` | `object` | да | Данные покупателя. |
| `receipt.client.name` | `string` | не указано явно | Имя или наименование покупателя. |
| `receipt.client.inn` | `string` | не указано явно | ИНН покупателя. |
| `receipt.client.email` | `string` | условно | Email покупателя. В объекте `receipt.client` обязательно одно из полей: `email` или `phone`. |
| `receipt.client.phone` | `string` | условно | Телефон покупателя. В объекте `receipt.client` обязательно одно из полей: `email` или `phone`. |
| `receipt.company` | `object` | да | Данные организации. |
| `receipt.company.email` | `string` | не указано явно | Email организации. |
| `receipt.company.sno` | `string` | не указано явно | Система налогообложения. |
| `receipt.company.inn` | `string` | да | ИНН организации. |
| `receipt.company.payment_address` | `string` | да | Адрес расчётов или сайт. |
| `receipt.company.location` | `string` | не указано явно | Место расчёта. |
| `receipt.agent_info` | `object` | не указано явно | Сведения об агенте на уровне чека. |
| `receipt.agent_info.type` | `string` | не указано явно | Тип агента. |
| `receipt.agent_info.paying_agent` | `object` | не указано явно | Данные платёжного агента. |
| `receipt.agent_info.paying_agent.operation` | `string` | не указано явно | Операция агента. |
| `receipt.agent_info.paying_agent.phones` | `array<string>` | не указано явно | Телефоны платёжного агента. |
| `receipt.agent_info.receive_payments_operator` | `object` | не указано явно | Данные оператора по приёму платежей. |
| `receipt.agent_info.receive_payments_operator.phones` | `array<string>` | не указано явно | Телефоны оператора по приёму платежей. |
| `receipt.agent_info.money_transfer_operator` | `object` | не указано явно | Данные оператора перевода денежных средств. |
| `receipt.agent_info.money_transfer_operator.phones` | `array<string>` | не указано явно | Телефоны оператора перевода денежных средств. |
| `receipt.agent_info.money_transfer_operator.name` | `string` | не указано явно | Наименование оператора перевода денежных средств. |
| `receipt.agent_info.money_transfer_operator.address` | `string` | не указано явно | Адрес оператора перевода денежных средств. |
| `receipt.agent_info.money_transfer_operator.inn` | `string` | не указано явно | ИНН оператора перевода денежных средств. |
| `receipt.supplier_info` | `object` | не указано явно | Сведения о поставщике на уровне чека. |
| `receipt.supplier_info.phones` | `array<string>` | не указано явно | Телефоны поставщика. |
| `receipt.supplier_info.name` | `string` | не указано явно | Наименование поставщика. |
| `receipt.supplier_info.inn` | `string` | не указано явно | ИНН поставщика. |
| `receipt.items` | `array<object>` | да | Позиции чека. |
| `receipt.items[].name` | `string` | да | Наименование позиции. |
| `receipt.items[].price` | `float` | да | Цена за единицу. |
| `receipt.items[].quantity` | `float` | да | Количество. |
| `receipt.items[].sum` | `float \| null` | да | Сумма по позиции. |
| `receipt.items[].measurement_unit` | `string` | не указано явно | Единица измерения. |
| `receipt.items[].payment_method` | `string` | не указано явно | Признак способа расчёта. |
| `receipt.items[].payment_object` | `string` | не указано явно | Признак предмета расчёта. |
| `receipt.items[].nomenclature_code` | `string` | не указано явно | Код товара или номенклатуры. |
| `receipt.items[].vat` | `object` | не указано явно | НДС по позиции. |
| `receipt.items[].vat.type` | `string` | да | Тип ставки НДС. |
| `receipt.items[].vat.sum` | `float \| null` | да | Сумма НДС. |
| `receipt.items[].agent_info` | `object` | не указано явно | Сведения об агенте на уровне позиции. |
| `receipt.items[].agent_info.type` | `string` | не указано явно | Тип агента. |
| `receipt.items[].agent_info.paying_agent` | `object` | не указано явно | Данные платёжного агента. |
| `receipt.items[].agent_info.paying_agent.operation` | `string` | не указано явно | Операция агента. |
| `receipt.items[].agent_info.paying_agent.phones` | `array<string>` | не указано явно | Телефоны платёжного агента. |
| `receipt.items[].agent_info.receive_payments_operator` | `object` | не указано явно | Данные оператора по приёму платежей. |
| `receipt.items[].agent_info.receive_payments_operator.phones` | `array<string>` | не указано явно | Телефоны оператора по приёму платежей. |
| `receipt.items[].agent_info.money_transfer_operator` | `object` | не указано явно | Данные оператора перевода денежных средств. |
| `receipt.items[].agent_info.money_transfer_operator.phones` | `array<string>` | не указано явно | Телефоны оператора перевода денежных средств. |
| `receipt.items[].agent_info.money_transfer_operator.name` | `string` | не указано явно | Наименование оператора перевода денежных средств. |
| `receipt.items[].agent_info.money_transfer_operator.address` | `string` | не указано явно | Адрес оператора перевода денежных средств. |
| `receipt.items[].agent_info.money_transfer_operator.inn` | `string` | не указано явно | ИНН оператора перевода денежных средств. |
| `receipt.items[].supplier_info` | `object` | не указано явно | Сведения о поставщике на уровне позиции. |
| `receipt.items[].supplier_info.phones` | `array<string>` | не указано явно | Телефоны поставщика. |
| `receipt.items[].supplier_info.name` | `string` | не указано явно | Наименование поставщика. |
| `receipt.items[].supplier_info.inn` | `string` | не указано явно | ИНН поставщика. |
| `receipt.items[].user_data` | `string` | не указано явно | Дополнительные пользовательские данные по позиции. |
| `receipt.items[].excise` | `float` | не указано явно | Сумма акциза. |
| `receipt.items[].country_code` | `string` | не указано явно | Цифровой код страны происхождения. |
| `receipt.items[].declaration_number` | `string` | не указано явно | Номер таможенной декларации. |
| `receipt.payments` | `array<object>` | да | Сведения об оплате. |
| `receipt.payments[].type` | `int` | да | Тип оплаты. |
| `receipt.payments[].sum` | `float \| null` | да | Сумма оплаты. |
| `receipt.vats` | `array<object>` | не указано явно | Итоговые НДС по чеку. |
| `receipt.vats[].type` | `string` | да | Тип ставки НДС. |
| `receipt.vats[].sum` | `float \| null` | да | Сумма НДС. |
| `receipt.total` | `float` | да | Общая сумма чека. |
| `receipt.additional_check_props` | `string` | не указано явно | Дополнительные реквизиты чека. |
| `receipt.cashier` | `string` | не указано явно | Кассир. |
| `receipt.additional_user_props` | `object` | не указано явно | Дополнительные пользовательские реквизиты. |
| `receipt.additional_user_props.name` | `string` | да | Имя дополнительного реквизита. |
| `receipt.additional_user_props.value` | `string` | да | Значение дополнительного реквизита. |
| `receipt.device_number` | `string` | не указано явно | Номер устройства. |
| `receipt.internet` | `bool` | не указано явно | Признак интернет-расчёта. |
| `receipt.cashless_payments` | `array<object>` | не указано явно | Дополнительные сведения по безналичным оплатам. |
| `receipt.cashless_payments[].sum` | `float` | да | Сумма безналичной оплаты. |
| `receipt.cashless_payments[].method` | `int` | да | Метод безналичной оплаты. |
| `receipt.cashless_payments[].id` | `string` | да | Идентификатор безналичной оплаты. |
| `receipt.cashless_payments[].additional_info` | `string` | не указано явно | Дополнительная информация по безналичной оплате. |

#### Успешный ответ

| field | type | required | description |
|---|---|---|---|
| `uuid` | `string` | да | Идентификатор зарегистрированного документа в системе АТОЛ. |
| `timestamp` | `string` | да | Время формирования ответа. |
| `error` | `object \| null` | да | Ошибка. В успешном ответе обычно `null`. |
| `status` | `string` | не указано явно | Статус обработки документа, например `wait`. |

#### Ответ с ошибкой

| field | type | required | description |
|---|---|---|---|
| `uuid` | `string` | не указано явно | Идентификатор документа, если уже был присвоен. |
| `error` | `object` | да | Объект ошибки. |
| `error.error_id` | `string` | да | Идентификатор ошибки. |
| `error.code` | `int` | да | Код ошибки. |
| `error.text` | `string` | да | Текст ошибки. |
| `error.type` | `string` | не указано явно | Тип ошибки. |
| `timestamp` | `string` | да | Время формирования ответа. |
| `status` | `string` | не указано явно | Статус обработки, например `fail`. |
| `group_code` | `string` | не указано явно | Код группы ККТ. |
| `daemon_code` | `string` | не указано явно | Код демона. |
| `device_code` | `string` | не указано явно | Код устройства. |
| `callback_url` | `string` | не указано явно | Callback URL, если был передан. |
| `payload` | `null` | не указано явно | Полезная нагрузка; в ошибочном ответе в тестах равна `null`. |

### 3.4 `POST /possystem/v4/{group_code}/buy_refund`

#### Тело запроса

Структура тела полностью совпадает с `sell_refund` и дублируется здесь намеренно.

| field | type | required | description |
|---|---|---|---|
| `timestamp` | `string` | да | Время формирования документа. |
| `source_id` | `int` | не указано явно | Числовой идентификатор источника. |
| `external_id` | `string` | да | Внешний идентификатор документа у интегратора. |
| `service` | `object` | не указано явно | Служебные параметры обработки. |
| `service.callback_url` | `string` | не указано явно | URL webhook, на который АТОЛ отправит результат обработки. |
| `receipt` | `object` | да | Тело чека возврата расхода. |
| `receipt.client` | `object` | да | Данные покупателя. |
| `receipt.client.name` | `string` | не указано явно | Имя или наименование покупателя. |
| `receipt.client.inn` | `string` | не указано явно | ИНН покупателя. |
| `receipt.client.email` | `string` | условно | Email покупателя. В объекте `receipt.client` обязательно одно из полей: `email` или `phone`. |
| `receipt.client.phone` | `string` | условно | Телефон покупателя. В объекте `receipt.client` обязательно одно из полей: `email` или `phone`. |
| `receipt.company` | `object` | да | Данные организации. |
| `receipt.company.email` | `string` | не указано явно | Email организации. |
| `receipt.company.sno` | `string` | не указано явно | Система налогообложения. |
| `receipt.company.inn` | `string` | да | ИНН организации. |
| `receipt.company.payment_address` | `string` | да | Адрес расчётов или сайт. |
| `receipt.company.location` | `string` | не указано явно | Место расчёта. |
| `receipt.agent_info` | `object` | не указано явно | Сведения об агенте на уровне чека. |
| `receipt.agent_info.type` | `string` | не указано явно | Тип агента. |
| `receipt.agent_info.paying_agent` | `object` | не указано явно | Данные платёжного агента. |
| `receipt.agent_info.paying_agent.operation` | `string` | не указано явно | Операция агента. |
| `receipt.agent_info.paying_agent.phones` | `array<string>` | не указано явно | Телефоны платёжного агента. |
| `receipt.agent_info.receive_payments_operator` | `object` | не указано явно | Данные оператора по приёму платежей. |
| `receipt.agent_info.receive_payments_operator.phones` | `array<string>` | не указано явно | Телефоны оператора по приёму платежей. |
| `receipt.agent_info.money_transfer_operator` | `object` | не указано явно | Данные оператора перевода денежных средств. |
| `receipt.agent_info.money_transfer_operator.phones` | `array<string>` | не указано явно | Телефоны оператора перевода денежных средств. |
| `receipt.agent_info.money_transfer_operator.name` | `string` | не указано явно | Наименование оператора перевода денежных средств. |
| `receipt.agent_info.money_transfer_operator.address` | `string` | не указано явно | Адрес оператора перевода денежных средств. |
| `receipt.agent_info.money_transfer_operator.inn` | `string` | не указано явно | ИНН оператора перевода денежных средств. |
| `receipt.supplier_info` | `object` | не указано явно | Сведения о поставщике на уровне чека. |
| `receipt.supplier_info.phones` | `array<string>` | не указано явно | Телефоны поставщика. |
| `receipt.supplier_info.name` | `string` | не указано явно | Наименование поставщика. |
| `receipt.supplier_info.inn` | `string` | не указано явно | ИНН поставщика. |
| `receipt.items` | `array<object>` | да | Позиции чека. |
| `receipt.items[].name` | `string` | да | Наименование позиции. |
| `receipt.items[].price` | `float` | да | Цена за единицу. |
| `receipt.items[].quantity` | `float` | да | Количество. |
| `receipt.items[].sum` | `float \| null` | да | Сумма по позиции. |
| `receipt.items[].measurement_unit` | `string` | не указано явно | Единица измерения. |
| `receipt.items[].payment_method` | `string` | не указано явно | Признак способа расчёта. |
| `receipt.items[].payment_object` | `string` | не указано явно | Признак предмета расчёта. |
| `receipt.items[].nomenclature_code` | `string` | не указано явно | Код товара или номенклатуры. |
| `receipt.items[].vat` | `object` | не указано явно | НДС по позиции. |
| `receipt.items[].vat.type` | `string` | да | Тип ставки НДС. |
| `receipt.items[].vat.sum` | `float \| null` | да | Сумма НДС. |
| `receipt.items[].agent_info` | `object` | не указано явно | Сведения об агенте на уровне позиции. |
| `receipt.items[].agent_info.type` | `string` | не указано явно | Тип агента. |
| `receipt.items[].agent_info.paying_agent` | `object` | не указано явно | Данные платёжного агента. |
| `receipt.items[].agent_info.paying_agent.operation` | `string` | не указано явно | Операция агента. |
| `receipt.items[].agent_info.paying_agent.phones` | `array<string>` | не указано явно | Телефоны платёжного агента. |
| `receipt.items[].agent_info.receive_payments_operator` | `object` | не указано явно | Данные оператора по приёму платежей. |
| `receipt.items[].agent_info.receive_payments_operator.phones` | `array<string>` | не указано явно | Телефоны оператора по приёму платежей. |
| `receipt.items[].agent_info.money_transfer_operator` | `object` | не указано явно | Данные оператора перевода денежных средств. |
| `receipt.items[].agent_info.money_transfer_operator.phones` | `array<string>` | не указано явно | Телефоны оператора перевода денежных средств. |
| `receipt.items[].agent_info.money_transfer_operator.name` | `string` | не указано явно | Наименование оператора перевода денежных средств. |
| `receipt.items[].agent_info.money_transfer_operator.address` | `string` | не указано явно | Адрес оператора перевода денежных средств. |
| `receipt.items[].agent_info.money_transfer_operator.inn` | `string` | не указано явно | ИНН оператора перевода денежных средств. |
| `receipt.items[].supplier_info` | `object` | не указано явно | Сведения о поставщике на уровне позиции. |
| `receipt.items[].supplier_info.phones` | `array<string>` | не указано явно | Телефоны поставщика. |
| `receipt.items[].supplier_info.name` | `string` | не указано явно | Наименование поставщика. |
| `receipt.items[].supplier_info.inn` | `string` | не указано явно | ИНН поставщика. |
| `receipt.items[].user_data` | `string` | не указано явно | Дополнительные пользовательские данные по позиции. |
| `receipt.items[].excise` | `float` | не указано явно | Сумма акциза. |
| `receipt.items[].country_code` | `string` | не указано явно | Цифровой код страны происхождения. |
| `receipt.items[].declaration_number` | `string` | не указано явно | Номер таможенной декларации. |
| `receipt.payments` | `array<object>` | да | Сведения об оплате. |
| `receipt.payments[].type` | `int` | да | Тип оплаты. |
| `receipt.payments[].sum` | `float \| null` | да | Сумма оплаты. |
| `receipt.vats` | `array<object>` | не указано явно | Итоговые НДС по чеку. |
| `receipt.vats[].type` | `string` | да | Тип ставки НДС. |
| `receipt.vats[].sum` | `float \| null` | да | Сумма НДС. |
| `receipt.total` | `float` | да | Общая сумма чека. |
| `receipt.additional_check_props` | `string` | не указано явно | Дополнительные реквизиты чека. |
| `receipt.cashier` | `string` | не указано явно | Кассир. |
| `receipt.additional_user_props` | `object` | не указано явно | Дополнительные пользовательские реквизиты. |
| `receipt.additional_user_props.name` | `string` | да | Имя дополнительного реквизита. |
| `receipt.additional_user_props.value` | `string` | да | Значение дополнительного реквизита. |
| `receipt.device_number` | `string` | не указано явно | Номер устройства. |
| `receipt.internet` | `bool` | не указано явно | Признак интернет-расчёта. |
| `receipt.cashless_payments` | `array<object>` | не указано явно | Дополнительные сведения по безналичным оплатам. |
| `receipt.cashless_payments[].sum` | `float` | да | Сумма безналичной оплаты. |
| `receipt.cashless_payments[].method` | `int` | да | Метод безналичной оплаты. |
| `receipt.cashless_payments[].id` | `string` | да | Идентификатор безналичной оплаты. |
| `receipt.cashless_payments[].additional_info` | `string` | не указано явно | Дополнительная информация по безналичной оплате. |

#### Успешный ответ

| field | type | required | description |
|---|---|---|---|
| `uuid` | `string` | да | Идентификатор зарегистрированного документа в системе АТОЛ. |
| `timestamp` | `string` | да | Время формирования ответа. |
| `error` | `object \| null` | да | Ошибка. В успешном ответе обычно `null`. |
| `status` | `string` | не указано явно | Статус обработки документа, например `wait`. |

#### Ответ с ошибкой

| field | type | required | description |
|---|---|---|---|
| `uuid` | `string` | не указано явно | Идентификатор документа, если уже был присвоен. |
| `error` | `object` | да | Объект ошибки. |
| `error.error_id` | `string` | да | Идентификатор ошибки. |
| `error.code` | `int` | да | Код ошибки. |
| `error.text` | `string` | да | Текст ошибки. |
| `error.type` | `string` | не указано явно | Тип ошибки. |
| `timestamp` | `string` | да | Время формирования ответа. |
| `status` | `string` | не указано явно | Статус обработки, например `fail`. |
| `group_code` | `string` | не указано явно | Код группы ККТ. |
| `daemon_code` | `string` | не указано явно | Код демона. |
| `device_code` | `string` | не указано явно | Код устройства. |
| `callback_url` | `string` | не указано явно | Callback URL, если был передан. |
| `payload` | `null` | не указано явно | Полезная нагрузка; в ошибочном ответе в тестах равна `null`. |

### 3.5 `POST /possystem/v4/{group_code}/sell_correction`

#### Тело запроса

| field | type | required | description |
|---|---|---|---|
| `timestamp` | `string` | да | Время формирования документа. |
| `source_id` | `int` | не указано явно | Числовой идентификатор источника. |
| `external_id` | `string` | да | Внешний идентификатор документа у интегратора. |
| `service` | `object` | не указано явно | Служебные параметры обработки. |
| `service.callback_url` | `string` | не указано явно | URL webhook, на который АТОЛ отправит результат обработки. |
| `correction` | `object` | да | Тело чека коррекции прихода. |
| `correction.company` | `object` | да | Данные организации. |
| `correction.company.sno` | `string` | не указано явно | Система налогообложения. |
| `correction.company.inn` | `string` | да | ИНН организации. |
| `correction.company.payment_address` | `string` | да | Адрес расчётов или сайт. |
| `correction.company.location` | `string` | не указано явно | Место расчёта. |
| `correction.client` | `object` | условно | Данные покупателя. Обязателен, если `correction.internet=true`. |
| `correction.client.email` | `string` | условно | Email покупателя. При `correction.internet=true` в объекте `correction.client` обязательно одно из полей: `email` или `phone`. |
| `correction.client.phone` | `string` | условно | Телефон покупателя. При `correction.internet=true` в объекте `correction.client` обязательно одно из полей: `email` или `phone`. |
| `correction.correction_info` | `object` | да | Основание и вид коррекции. |
| `correction.correction_info.type` | `string` | да | Тип коррекции. |
| `correction.correction_info.base_date` | `string` | да | Дата документа-основания. |
| `correction.correction_info.base_number` | `string` | не указано явно | Номер документа-основания. В PDF поле выглядит обязательным, но на тестовом стенде `sell_correction` успешно обработан без этого поля. |
| `correction.payments` | `array<object>` | да | Сведения об оплате. |
| `correction.payments[].type` | `int` | да | Тип оплаты. |
| `correction.payments[].sum` | `float \| null` | да | Сумма оплаты. |
| `correction.vats` | `array<object>` | да | НДС по коррекции. |
| `correction.vats[].type` | `string` | да | Тип ставки НДС. |
| `correction.vats[].sum` | `float \| null` | да | Сумма НДС. |
| `correction.cashier` | `string` | не указано явно | Кассир. |
| `correction.device_number` | `string` | не указано явно | Номер устройства. |
| `correction.internet` | `bool` | не указано явно | Признак интернет-расчёта. |

#### Успешный ответ

| field | type | required | description |
|---|---|---|---|
| `uuid` | `string` | да | Идентификатор зарегистрированного документа в системе АТОЛ. |
| `timestamp` | `string` | да | Время формирования ответа. |
| `error` | `object \| null` | да | Ошибка. В успешном ответе обычно `null`. |
| `status` | `string` | не указано явно | Статус обработки документа, например `wait`. |

#### Ответ с ошибкой

| field | type | required | description |
|---|---|---|---|
| `uuid` | `string` | не указано явно | Идентификатор документа, если уже был присвоен. |
| `error` | `object` | да | Объект ошибки. |
| `error.error_id` | `string` | да | Идентификатор ошибки. |
| `error.code` | `int` | да | Код ошибки. |
| `error.text` | `string` | да | Текст ошибки. |
| `error.type` | `string` | не указано явно | Тип ошибки. |
| `timestamp` | `string` | да | Время формирования ответа. |
| `status` | `string` | не указано явно | Статус обработки, например `fail`. |
| `group_code` | `string` | не указано явно | Код группы ККТ. |
| `daemon_code` | `string` | не указано явно | Код демона. |
| `device_code` | `string` | не указано явно | Код устройства. |
| `callback_url` | `string` | не указано явно | Callback URL, если был передан. |
| `payload` | `null` | не указано явно | Полезная нагрузка; в ошибочном ответе в тестах равна `null`. |

### 3.6 `POST /possystem/v4/{group_code}/buy_correction`

#### Тело запроса

Структура тела полностью совпадает с `sell_correction` и дублируется здесь намеренно.

| field | type | required | description |
|---|---|---|---|
| `timestamp` | `string` | да | Время формирования документа. |
| `source_id` | `int` | не указано явно | Числовой идентификатор источника. |
| `external_id` | `string` | да | Внешний идентификатор документа у интегратора. |
| `service` | `object` | не указано явно | Служебные параметры обработки. |
| `service.callback_url` | `string` | не указано явно | URL webhook, на который АТОЛ отправит результат обработки. |
| `correction` | `object` | да | Тело чека коррекции расхода. |
| `correction.company` | `object` | да | Данные организации. |
| `correction.company.sno` | `string` | не указано явно | Система налогообложения. |
| `correction.company.inn` | `string` | да | ИНН организации. |
| `correction.company.payment_address` | `string` | да | Адрес расчётов или сайт. |
| `correction.company.location` | `string` | не указано явно | Место расчёта. |
| `correction.client` | `object` | условно | Данные покупателя. Обязателен, если `correction.internet=true`. |
| `correction.client.email` | `string` | условно | Email покупателя. При `correction.internet=true` в объекте `correction.client` обязательно одно из полей: `email` или `phone`. |
| `correction.client.phone` | `string` | условно | Телефон покупателя. При `correction.internet=true` в объекте `correction.client` обязательно одно из полей: `email` или `phone`. |
| `correction.correction_info` | `object` | да | Основание и вид коррекции. |
| `correction.correction_info.type` | `string` | да | Тип коррекции. |
| `correction.correction_info.base_date` | `string` | да | Дата документа-основания. |
| `correction.correction_info.base_number` | `string` | не указано явно | Номер документа-основания. В PDF поле выглядит обязательным, но на тестовом стенде `buy_correction`/`sell_correction` следует считать его фактически необязательным до дополнительной проверки. |
| `correction.payments` | `array<object>` | да | Сведения об оплате. |
| `correction.payments[].type` | `int` | да | Тип оплаты. |
| `correction.payments[].sum` | `float \| null` | да | Сумма оплаты. |
| `correction.vats` | `array<object>` | да | НДС по коррекции. |
| `correction.vats[].type` | `string` | да | Тип ставки НДС. |
| `correction.vats[].sum` | `float \| null` | да | Сумма НДС. |
| `correction.cashier` | `string` | не указано явно | Кассир. |
| `correction.device_number` | `string` | не указано явно | Номер устройства. |
| `correction.internet` | `bool` | не указано явно | Признак интернет-расчёта. |

#### Успешный ответ

| field | type | required | description |
|---|---|---|---|
| `uuid` | `string` | да | Идентификатор зарегистрированного документа в системе АТОЛ. |
| `timestamp` | `string` | да | Время формирования ответа. |
| `error` | `object \| null` | да | Ошибка. В успешном ответе обычно `null`. |
| `status` | `string` | не указано явно | Статус обработки документа, например `wait`. |

#### Ответ с ошибкой

| field | type | required | description |
|---|---|---|---|
| `uuid` | `string` | не указано явно | Идентификатор документа, если уже был присвоен. |
| `error` | `object` | да | Объект ошибки. |
| `error.error_id` | `string` | да | Идентификатор ошибки. |
| `error.code` | `int` | да | Код ошибки. |
| `error.text` | `string` | да | Текст ошибки. |
| `error.type` | `string` | не указано явно | Тип ошибки. |
| `timestamp` | `string` | да | Время формирования ответа. |
| `status` | `string` | не указано явно | Статус обработки, например `fail`. |
| `group_code` | `string` | не указано явно | Код группы ККТ. |
| `daemon_code` | `string` | не указано явно | Код демона. |
| `device_code` | `string` | не указано явно | Код устройства. |
| `callback_url` | `string` | не указано явно | Callback URL, если был передан. |
| `payload` | `null` | не указано явно | Полезная нагрузка; в ошибочном ответе в тестах равна `null`. |

## 4. `GET /possystem/v4/{group_code}/report/{uuid}`

Полный URL:

```text
https://online.atol.ru/possystem/v4/{group_code}/report/{uuid}
```

Дополнительно в PDF показан вариант:

```text
https://online.atol.ru/possystem/v4/{group_code}/report/{uuid}?token=<token>
```

### Параметры URL

| field | type | required | description |
|---|---|---|---|
| `group_code` | `string` | да | Код группы ККТ. |
| `uuid` | `string` | да | Идентификатор ранее зарегистрированного документа. |
| `token` | `string` | не указано явно | Токен в query string, если используется такой вариант авторизации. |

### Тело запроса

Отсутствует.

### Успешный ответ

| field | type | required | description |
|---|---|---|---|
| `uuid` | `string` | да | Идентификатор документа. |
| `error` | `object \| null` | да | Ошибка. В успешном ответе обычно `null`. |
| `status` | `string` | не указано явно | Статус обработки документа, например `done`. |
| `payload` | `object` | да | Фискальные реквизиты документа после успешной обработки. |
| `payload.total` | `float` | да | Итоговая сумма документа. |
| `payload.fns_site` | `string` | да | Адрес сайта ФНС. |
| `payload.fn_number` | `string` | да | Номер ФН. |
| `payload.shift_number` | `int` | да | Номер смены. |
| `payload.receipt_datetime` | `string` | да | Дата и время документа. |
| `payload.fiscal_receipt_number` | `int` | да | Номер чека в смене. |
| `payload.fiscal_document_number` | `int` | да | Номер фискального документа. |
| `payload.ecr_registration_number` | `string` | да | Регистрационный номер ККТ. |
| `payload.fiscal_document_attribute` | `int` | да | Фискальный признак документа. |
| `payload.ofd_inn` | `string` | не указано явно | ИНН ОФД. |
| `payload.ofd_receipt_url` | `string` | не указано явно | Ссылка на чек у ОФД. |
| `timestamp` | `string` | да | Время формирования ответа. |
| `group_code` | `string` | да | Код группы ККТ. |
| `daemon_code` | `string` | да | Код демона. |
| `device_code` | `string` | да | Код устройства. |
| `external_id` | `string` | не указано явно | Внешний идентификатор документа у интегратора. |
| `callback_url` | `string` | не указано явно | Callback URL, переданный при регистрации, либо пустая строка. |
| `warnings` | `object` | не указано явно | Предупреждения по доставке результата. |
| `warnings.callback_url` | `string` | не указано явно | Предупреждение по callback URL. |

### Ответ с ошибкой

| field | type | required | description |
|---|---|---|---|
| `uuid` | `string` | не указано явно | Идентификатор документа. |
| `error` | `object` | да | Объект ошибки. |
| `error.error_id` | `string` | да | Идентификатор ошибки. |
| `error.code` | `int` | да | Код ошибки. |
| `error.text` | `string` | да | Текст ошибки. |
| `error.type` | `string` | не указано явно | Тип ошибки. |
| `timestamp` | `string` | да | Время формирования ответа. |
| `status` | `string` | не указано явно | Статус обработки, если сервис его возвращает. |
| `group_code` | `string` | не указано явно | Код группы ККТ. |
| `daemon_code` | `string` | не указано явно | Код демона. |
| `device_code` | `string` | не указано явно | Код устройства. |
| `callback_url` | `string` | не указано явно | Callback URL, если он был задан. |
| `payload` | `null` | не указано явно | Полезная нагрузка в error-ответе. |

## 5. `POST <callback_url>`

Это не endpoint АТОЛ, а webhook на стороне клиента, который сервис АТОЛ вызывает для доставки результата обработки.

### Параметры URL

URL целиком задаётся интегратором в `service.callback_url`.

### Тело входящего запроса

По составу полей webhook повторяет результат `report` и дублируется здесь намеренно.

| field | type | required | description |
|---|---|---|---|
| `uuid` | `string` | да | Идентификатор документа. |
| `error` | `object \| null` | да | Ошибка. Для успешной доставки результата обычно `null`. |
| `error.error_id` | `string` | не указано явно | Идентификатор ошибки. |
| `error.code` | `int` | не указано явно | Код ошибки. |
| `error.text` | `string` | не указано явно | Текст ошибки. |
| `error.type` | `string` | не указано явно | Тип ошибки. |
| `status` | `string` | не указано явно | Статус обработки документа. |
| `payload` | `object` | да | Фискальные реквизиты документа после успешной обработки. |
| `payload.total` | `float` | да | Итоговая сумма документа. |
| `payload.fns_site` | `string` | да | Адрес сайта ФНС. |
| `payload.fn_number` | `string` | да | Номер ФН. |
| `payload.shift_number` | `int` | да | Номер смены. |
| `payload.receipt_datetime` | `string` | да | Дата и время документа. |
| `payload.fiscal_receipt_number` | `int` | да | Номер чека в смене. |
| `payload.fiscal_document_number` | `int` | да | Номер фискального документа. |
| `payload.ecr_registration_number` | `string` | да | Регистрационный номер ККТ. |
| `payload.fiscal_document_attribute` | `int` | да | Фискальный признак документа. |
| `payload.ofd_inn` | `string` | не указано явно | ИНН ОФД. |
| `payload.ofd_receipt_url` | `string` | не указано явно | Ссылка на чек у ОФД. |
| `timestamp` | `string` | да | Время формирования уведомления. |
| `group_code` | `string` | да | Код группы ККТ. |
| `daemon_code` | `string` | да | Код демона. |
| `device_code` | `string` | да | Код устройства. |
| `external_id` | `string` | не указано явно | Внешний идентификатор документа у интегратора. |
| `callback_url` | `string` | не указано явно | Callback URL, на который отправлено уведомление. |
| `warnings` | `object` | не указано явно | Предупреждения по доставке результата. |
| `warnings.callback_url` | `string` | не указано явно | Предупреждение по callback URL. |

### Ответ сервиса клиента на webhook

В PDF, доступном локально для проверки, структура ответа клиента на webhook не выделена как отдельная JSON-схема. Обычно ожидается успешный HTTP-статус от принимающей стороны.

## 6. Проверка на тестовом стенде

Ниже перечислены только те наблюдения, которые подтверждены живыми запросами 13.03.2026.

### 6.1 `POST /possystem/v4/getToken`

Проверенный минимальный запрос:

```json
{
  "login": "v4-online-atol-ru",
  "pass": "iGFFuihss"
}
```

Фактически полученный успешный ответ:

| field | type | required | description |
|---|---|---|---|
| `token` | `string` | да | Токен доступа. |
| `error` | `null` | да | Ошибка отсутствует. |
| `timestamp` | `string` | да | Время формирования ответа. |

### 6.2 `POST /possystem/v4/{group_code}/sell`

Проверенный минимальный запрос, успешно прошедший до статуса `done`:

```json
{
  "timestamp": "2026-03-13T10:20:30+03:00",
  "external_id": "codex-sell-20260313-102030",
  "receipt": {
    "client": {
      "email": "test@example.com"
    },
    "company": {
      "email": "test@example.com",
      "sno": "osn",
      "inn": "5544332219",
      "payment_address": "https://v4.online.atol.ru"
    },
    "items": [
      {
        "name": "Test item",
        "price": 100,
        "quantity": 1,
        "sum": 100,
        "vat": {
          "type": "none",
          "sum": 0
        }
      }
    ],
    "payments": [
      {
        "type": 1,
        "sum": 100
      }
    ],
    "total": 100
  }
}
```

Проверенные выводы:

- `receipt.client` реально обязателен: запрос без него вернул `PropertyRequired: #/receipt.client`.
- для минимального рабочего запроса достаточно `receipt.client.email`; `phone` не потребовался.
- запрос с неконсистентным `vat` был принят в очередь, но в `report` завершился ошибкой бизнес-валидации `Некорректная сумма налога`.

Фактически полученный успешный ответ на `POST`:

| field | type | required | description |
|---|---|---|---|
| `uuid` | `string` | да | Идентификатор документа. |
| `status` | `string` | да | На тестовом стенде возвращён `wait`. |
| `error` | `null` | да | Ошибка отсутствует. |
| `timestamp` | `string` | да | Время формирования ответа. |

Фактически полученный ответ на ошибку валидации входящего `POST`:

| field | type | required | description |
|---|---|---|---|
| `status` | `string` | да | На тестовом стенде возвращён `fail`. |
| `error.error_id` | `string` | да | Идентификатор ошибки. |
| `error.code` | `int` | да | Код ошибки. |
| `error.text` | `string` | да | Текст ошибки. |
| `error.type` | `string` | да | Тип ошибки. |
| `timestamp` | `string` | да | Время формирования ответа. |

### 6.3 `GET /possystem/v4/{group_code}/report/{uuid}`

Фактически полученный успешный ответ:

| field | type | required | description |
|---|---|---|---|
| `callback_url` | `string` | да | На тестовом стенде пришла пустая строка. |
| `daemon_code` | `string` | да | Код демона. |
| `device_code` | `string` | да | Код устройства. |
| `warnings` | `null` | да | На тестовом стенде пришёл `null`. |
| `error` | `null` | да | Ошибка отсутствует. |
| `external_id` | `string` | да | Внешний идентификатор документа. |
| `group_code` | `string` | да | Код группы ККТ. |
| `payload` | `object` | да | Фискальные реквизиты документа. |
| `payload.ecr_registration_number` | `string` | да | Регистрационный номер ККТ. |
| `payload.fiscal_document_attribute` | `int` | да | Фискальный признак документа. |
| `payload.fiscal_document_number` | `int` | да | Номер фискального документа. |
| `payload.fiscal_receipt_number` | `int` | да | Номер чека в смене. |
| `payload.fn_number` | `string` | да | Номер ФН. |
| `payload.fns_site` | `string` | да | Сайт ФНС. |
| `payload.receipt_datetime` | `string` | да | Дата и время документа. |
| `payload.shift_number` | `int` | да | Номер смены. |
| `payload.total` | `float` | да | Итоговая сумма. |
| `payload.ofd_inn` | `string` | да | ИНН ОФД. |
| `payload.ofd_receipt_url` | `string` | да | Ссылка на чек у ОФД. |
| `status` | `string` | да | На тестовом стенде возвращён `done`. |
| `uuid` | `string` | да | Идентификатор документа. |
| `timestamp` | `string` | да | Время формирования ответа. |

Фактически полученный ответ `report` при ошибке обработки документа:

| field | type | required | description |
|---|---|---|---|
| `callback_url` | `string` | да | На тестовом стенде пришла пустая строка. |
| `daemon_code` | `string` | да | Код демона. |
| `device_code` | `string` | да | Код устройства. |
| `warnings` | `null` | да | На тестовом стенде пришёл `null`. |
| `error.code` | `int` | да | Код ошибки. |
| `error.text` | `string` | да | Текст ошибки. |
| `error.type` | `string` | да | Тип ошибки. |
| `external_id` | `string` | да | Внешний идентификатор документа. |
| `group_code` | `string` | да | Код группы ККТ. |
| `payload` | `null` | да | Полезная нагрузка отсутствует. |
| `status` | `string` | да | На тестовом стенде возвращён `fail`. |
| `uuid` | `string` | да | Идентификатор документа. |
| `timestamp` | `string` | да | Время формирования ответа. |

### 6.4 `POST /possystem/v4/{group_code}/sell_correction`

Проверенный минимальный запрос, успешно прошедший до статуса `done`:

```json
{
  "timestamp": "2026-03-13T10:21:10+03:00",
  "external_id": "codex-corr-20260313-102110",
  "correction": {
    "company": {
      "sno": "osn",
      "inn": "5544332219",
      "payment_address": "https://v4.online.atol.ru"
    },
    "correction_info": {
      "type": "self",
      "base_date": "13.03.2026",
      "base_number": "1"
    },
    "payments": [
      {
        "type": 1,
        "sum": 100
      }
    ],
    "vats": [
      {
        "type": "none",
        "sum": 0
      }
    ]
  }
}
```

Проверенные выводы:

- `correction.client` не требуется, пока `correction.internet` не установлен в `true`.
- при `correction.internet=true` без `correction.client` стенд вернул ошибку валидации с `PropertyRequired: #/correction.client`.
- `correction.correction_info.base_number` на тестовом стенде не обязателен: запрос без него был успешно обработан со статусом `done`.

Фактически полученный успешный ответ на `POST` совпал по составу полей с успешным `POST /sell`:

| field | type | required | description |
|---|---|---|---|
| `uuid` | `string` | да | Идентификатор документа. |
| `status` | `string` | да | На тестовом стенде возвращён `wait`. |
| `error` | `null` | да | Ошибка отсутствует. |
| `timestamp` | `string` | да | Время формирования ответа. |
