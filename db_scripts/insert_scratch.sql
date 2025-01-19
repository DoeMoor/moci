
INSERT INTO manufacturers ("name")
VALUES ('Global Electronics'),
       ('Intern'),
       ('manufacturer1')
;
-- RETURNING *;
-- DELETE FROM manufacturers;

INSERT INTO mini_pcie_modules_type ("name", module_type_number)
VALUES ('4G/GPS SIMCOM 7600-G', '30100201'),
       ('4G/GPS SIMCOM 7600-G', '30100202');

-- RETURNING *;
-- DELETE FROM mini_pcie_modules_type;

insert into enclosure (serial_number, controller_types_id, manufacturers_id)
values ('ineni3423',
         (select id from controller_types limit 1),
         (select id from manufacturers limit 1)),
        ('inenssi3423',
         (select id from controller_types limit 1 offset 1),
         (select manufacturers.id from manufacturers limit 1 offset 1));

-- update enclosure set controller_id = (select id from controllers limit 1) where serial_number = 'ineni3423';

INSERT INTO mini_pcie_modules (mini_pcie_modules_type_id, serial_number)
VALUES ((SELECT id FROM mini_pcie_modules_type LIMIT 1),
        'MP062141639F5A4'),
       ((SELECT id FROM mini_pcie_modules_type LIMIT 1),
        'MP06214163925A4')
;
-- RETURNING *;
-- DELETE FROM mini_pcie_modules;


INSERT INTO m2_modules_types ("name", module_type_number)
VALUES ('WiFi/BT Murata 1MW', '30100101'),
       ('WiFi/BT Murata 1DX', '30100102')
;
-- RETURNING *;
-- DELETE FROM m2_modules_types;


INSERT INTO customers ("name")
VALUES ('H2Consultancy'),
       ('ARGO-ANLEG')
;
-- RETURNING *;
-- DELETE FROM customers;


INSERT INTO projects ("name")
VALUES ('B-22-01-07 PAUL'),
       ('Water control'),
       ('B-23-10-11 Volvo')
;
-- RETURNING *;
-- DELETE FROM projects;



insert into display_adapters_type (name)
values ('av123z7m'),
       ('av123z74');


INSERT INTO display_adapters (display_adapters_type_name, manufacturer_qr_code)
VALUES ('av123z7m', 'GE00000626856'),
       ('av123z74', 'GE00000623856')
;
-- RETURNING *;
-- DELETE FROM display_adapters;


INSERT INTO orders (customers_id)
VALUES ((SELECT id FROM customers LIMIT 1))
;
-- RETURNING *;
-- DELETE FROM orders;


INSERT INTO can_termination_confs
(can_1_terminated,
 can_2_terminated,
 can_3_terminated,
 can_4_terminated)
VALUES ('true', 'true', 'true', 'true')
;
-- RETURNING *;
-- DELETE FROM can_termination_confs ;

INSERT INTO controller_types
    ("name", io_module_slot_amount, amount_of_connectors, slot_pinout_json)
VALUES ('moduline 4', '8', '5',
        '{
          "controller": "moduline 4",
          "connector_A": {
            "pin_count": "26",
            "pin": {
              "slot_1": {
                "4": "odd_13",
                "5": "odd_3",
                "6": "odd_2",
                "7": "odd_1",
                "11": "odd_6",
                "12": "odd_5",
                "13": "odd_4",
                "17": "odd_9",
                "18": "odd_8",
                "19": "odd_7",
                "24": "odd_12",
                "25": "odd_11",
                "26": "odd_10"
              },
              "slot_2": {
                "1": "even_1",
                "2": "even_2",
                "3": "even_3",
                "8": "even_4",
                "9": "even_5",
                "10": "even_6",
                "14": "even_7",
                "15": "even_8",
                "16": "even_9",
                "20": "even_10",
                "21": "even_11",
                "22": "even_12",
                "23": "even_13"
              }
            }
          },
          "connector_B": {
            "pin_count": "26",
            "pin": {
              "slot_3": {
                "4": "odd_13",
                "5": "odd_3",
                "6": "odd_2",
                "7": "odd_1",
                "11": "odd_6",
                "12": "odd_5",
                "13": "odd_4",
                "17": "odd_9",
                "18": "odd_8",
                "19": "odd_7",
                "24": "odd_12",
                "25": "odd_11",
                "26": "odd_10"
              },
              "slot_4": {
                "1": "even_1",
                "2": "even_2",
                "3": "even_3",
                "8": "even_4",
                "9": "even_5",
                "10": "even_6",
                "14": "even_7",
                "15": "even_8",
                "16": "even_9",
                "20": "even_10",
                "21": "even_11",
                "22": "even_12",
                "23": "even_13"
              }
            }
          },
          "connector_C": {
            "pin_count": "26",
            "pin": {
              "system": {
                "1": "Reset",
                "2": "K15a",
                "3": "K15b",
                "4": "K15c",
                "5": "K30",
                "6": "K30",
                "7": "K30",
                "8": "LIN com",
                "9": "LIN sup",
                "10": "LIN gnd",
                "11": "CAN 4 gnd",
                "12": "CAN 4 high",
                "13": "CAN 4 low",
                "14": "CAN 1 low",
                "15": "CAN 1 high",
                "16": "CAN 1 gnd",
                "17": "CAN 3 gnd",
                "18": "CAN 3 high",
                "19": "CAN 3 low",
                "20": "CAN 2 low",
                "21": "CAN 2 high",
                "22": "CAN 2 gnd",
                "23": "?",
                "24": "K31",
                "25": "K31",
                "26": "K31"
              }
            }
          },
          "connector_D": {
            "pin_count": "26",
            "pin": {
              "slot_5": {
                "4": "odd_13",
                "5": "odd_3",
                "6": "odd_2",
                "7": "odd_1",
                "11": "odd_6",
                "12": "odd_5",
                "13": "odd_4",
                "17": "odd_9",
                "18": "odd_8",
                "19": "odd_7",
                "24": "odd_12",
                "25": "odd_11",
                "26": "odd_10"
              },
              "slot_6": {
                "1": "even_1",
                "2": "even_2",
                "3": "even_3",
                "8": "even_4",
                "9": "even_5",
                "10": "even_6",
                "14": "even_7",
                "15": "even_8",
                "16": "even_9",
                "20": "even_10",
                "21": "even_11",
                "22": "even_12",
                "23": "even_13"
              }
            }
          },
          "connector_E": {
            "pin_count": "26",
            "pin": {
              "slot_7": {
                "4": "odd_13",
                "5": "odd_3",
                "6": "odd_2",
                "7": "odd_1",
                "11": "odd_6",
                "12": "odd_5",
                "13": "odd_4",
                "17": "odd_9",
                "18": "odd_8",
                "19": "odd_7",
                "24": "odd_12",
                "25": "odd_11",
                "26": "odd_10"
              },
              "slot_8": {
                "1": "even_1",
                "2": "even_2",
                "3": "even_3",
                "8": "even_4",
                "9": "even_5",
                "10": "even_6",
                "14": "even_7",
                "15": "even_8",
                "16": "even_9",
                "20": "even_10",
                "21": "even_11",
                "22": "even_12",
                "23": "even_13"
              }
            }
          }
        }'),
       ('moduline mini', '4', '2',
        '{
          "controller": "moduline mini",
          "connector_A": {
            "pin_count": "34",
            "pin": {
              "slot_1": {
                "4": "odd_13",
                "5": "odd_3",
                "6": "odd_2",
                "7": "odd_1",
                "13": "odd_6",
                "14": "odd_5",
                "15": "odd_4",
                "21": "odd_9",
                "22": "odd_8",
                "23": "odd_7",
                "30": "odd_12",
                "31": "odd_11",
                "32": "odd_10"
              },
              "slot_2": {
                "1": "even_1",
                "2": "even_2",
                "3": "even_3",
                "10": "even_4",
                "11": "even_5",
                "12": "even_6",
                "18": "even_7",
                "19": "even_8",
                "20": "even_9",
                "26": "even_10",
                "27": "even_11",
                "28": "even_12",
                "29": "even_13"
              },
              "system": {
                "8": "nc",
                "9": "nc",
                "16": "K15b",
                "17": "K15c",
                "24": "CAN 2 low",
                "25": "CAN 2 high",
                "33": "nc",
                "34": "nc"
              }
            }
          },
          "connector_B": {
            "pin_count": "34",
            "pin": {
              "slot_1": {
                "4": "odd_13",
                "5": "odd_3",
                "6": "odd_2",
                "7": "odd_1",
                "13": "odd_6",
                "14": "odd_5",
                "15": "odd_4",
                "21": "odd_9",
                "22": "odd_8",
                "23": "odd_7",
                "30": "odd_12",
                "31": "odd_11",
                "32": "odd_10"
              },
              "slot_2": {
                "1": "even_1",
                "2": "even_2",
                "3": "even_3",
                "10": "even_4",
                "11": "even_5",
                "12": "even_6",
                "18": "even_7",
                "19": "even_8",
                "20": "even_9",
                "26": "even_10",
                "27": "even_11",
                "28": "even_12",
                "29": "even_13"
              },
              "system": {
                "8": "K30",
                "9": "K30",
                "16": "Reset",
                "17": "K15a",
                "24": "CAN 1 low",
                "25": "CAN 1 high",
                "33": "K31",
                "34": "K31"
              }
            }
          }
        }');


INSERT INTO controller_pcb_hw_versions_rev (revision)
VALUES ('a'),
       ('b')
;
-- RETURNING *;
-- DELETE FROM controller_pcb_hw_version_revisions ;


INSERT INTO controller_pcb_hw_versions (version_number, revision)
VALUES ('306', (SELECT revision FROM controller_pcb_hw_versions_rev LIMIT 1)),
       ('306', NULL),
       ('307', NULL)
;
-- RETURNING *;
-- DELETE FROM controller_pcb_hw_versions;


INSERT INTO io_module_types
(module_type_name,
 module_type_number,
 pinout_json)
VALUES ('Input module 10 channel',
        '201002',
        '{
          "module_type_name": "input",
          "module_type_number": 201002,
          "channels": 10,
          "pin": {
            "1": "Supply 1",
            "2": "Supply 2",
            "3": "Channel 9",
            "4": "Channel 1",
            "5": "Channel 3",
            "6": "Channel 5",
            "7": "Channel 2",
            "8": "Channel 4",
            "9": "Channel 6",
            "10": "Sensor ground",
            "11": "Channel 7",
            "12": "Channel 8",
            "13": "Channel 10"
          }
        }'),
       ('Input module 10 channel',
        '201003',
        '{
          "module_type_name": "input",
          "module_type_number": 201003,
          "channels": 10,
          "pin": {
            "1": "Channel 1",
            "2": "Channel 4",
            "3": "Channel 7",
            "4": "Channel 2",
            "5": "Channel 5",
            "6": "Channel 8",
            "7": "Channel 3",
            "8": "Channel 6",
            "9": "Channel 9",
            "10": "Sensor ground",
            "11": "Sensor ground",
            "12": "Sensor ground",
            "13": "Channel 10"
          }
        }')
;
-- RETURNING*;
-- DELETE FROM io_module_types ;

INSERT INTO io_module_hw_versions (hw_version)
VALUES ('02'),
       ('03'),
       ('04')
;
-- RETURNING*;
-- DELETE FROM io_module_hw_versions ;


INSERT INTO io_modules
(io_module_types_id,
 manufacturer_top_qr_code,
 manufacturer_bottom_qr_code,
 rma_number,
 io_module_hw_versions_id,
 order_id)
VALUES ((select id FROM io_module_types LIMIT 1),
        'GE00000623365',
        'GE00000623033',
        'rma_number',
        (select id FROM io_module_hw_versions LIMIT 1),
        (select id from orders limit 1)),
       ((select id FROM io_module_types LIMIT 1 OFFSET 1),
        'GE00000623367',
        'GE00000623027',
        'rma_number',
        (select id FROM io_module_hw_versions LIMIT 1), null),
       ((select id FROM io_module_types LIMIT 1 OFFSET 1),
        'GE00000523367',
        'GE00000423027',
        'rma_number',
        (select id FROM io_module_hw_versions LIMIT 1 OFFSET 1), null)
;
-- RETURNING*;
-- DELETE FROM io_modules ;


INSERT INTO CONTROLLERS
(CONTROLLER_TYPES_ID,
 PROJECTS_ID,
 DESCRIPTION,
 controllers_pcb_hw_versions_id,
 PCB_VERSION_NUMBER,
 SERIAL_NUMBER,
 MANUFACTURERS_ID,
 MAC_ADDRESS,
 SIM_NUMBER,
 MINI_PCIE_MODULES_ID,
 m2_modules_types_id,
 ARTICLE_NUMBER,
 QR_CODE,
 CAN_TERMINATION_CONFS_ID,
 USB,
 serial,
 manufacturer_qr_code,
 ORDER_ID)
VALUES ((SELECT id FROM controller_types LIMIT 1),
        (SELECT id FROM projects LIMIT 1),
        'A-22-03-16',
        (SELECT id FROM controller_pcb_hw_versions LIMIT 1),
        '25',
        'A4CG-B01A-A050-A001',
        (SELECT id FROM manufacturers LIMIT 1),
        '00:0c:c6:89:d0:2f',
        '89883070000012511109',
        (SELECT id FROM mini_pcie_modules LIMIT 1),
        (SELECT id FROM m2_modules_types LIMIT 1),
        '29040007',
        '{"s":"A4CG-B01A-A050-A001","b":"00:0c:c6:89:d0:2f","c":"29040007"}',
        (SELECT id FROM can_termination_confs LIMIT 1),
        'true',
        'false',
        'GE00000626856',
        (SELECT id FROM orders LIMIT 1)),
       ((SELECT id FROM controller_types LIMIT 1 OFFSET 1),
        (SELECT id FROM projects LIMIT 1 OFFSET 1),
        'A-22-03-15',
        (SELECT id FROM controller_pcb_hw_versions LIMIT 1 OFFSET 1),
        '23',
        'A4CG-B01A-A050-A004',
        (SELECT id FROM manufacturers LIMIT 1),
        '00:0c:c6:89:d4:2f',
        '89883070000013521109',
        (SELECT id FROM mini_pcie_modules LIMIT 1 OFFSET 1),
        (SELECT id FROM m2_modules_types LIMIT 1 OFFSET 1),
        '2904f0017',
        '{"s":"A4CG-B01A-A250-A001","b":"00:0c:c6:89:d0:2f","c":"29040007"}',
        (SELECT id FROM can_termination_confs LIMIT 1),
        'true',
        'false',
        'GE00300f626856',
        (SELECT id FROM orders LIMIT 1))
;
-- RETURNING *;
-- DELETE FROM CONTROLLERS;

INSERT INTO led_daughter_board
(manufacturer_qr_code,
 controllers_id)
VALUES ('GE000006263835',
        (select id FROM controllers LIMIT 1))
;
---reTURNING*;
---DELETE FROM led_daughter_board ;

insert into enclosure (serial_number, controller_types_id, manufacturers_id)
values ('rstar323ine2323',
        (select id from controller_types limit 1),
        (select id from manufacturers limit 1)),
       ('rstar323wfine2323',
        (select id from controller_types limit 1 offset 1),
        (select id from manufacturers limit 1 offset 1));

INSERT INTO controller_slot
(controllers_id,
 io_modules_id,
 slot_number)
VALUES ((select id FROM controllers limit 1), (select id FROM io_modules LIMIT 1), '1'),
       ((select id FROM controllers limit 1), (select id FROM io_modules LIMIT 1 offset 1), '2'),
       ((select id FROM controllers limit 1 OFFSET 1), (select id FROM io_modules LIMIT 1 offset 2), '1')
;
-- RETURNING *;
-- DELETE FROM controller_slot;


INSERT INTO customer_projects_junction (customers_id, projects_id)
VALUES ((select id FROM customers limit 1), (select id FROM projects LIMIT 1))
;
-- RETURNING *;
-- DELETE FROM customer_projects_junction;







