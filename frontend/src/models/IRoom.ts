import { EmployeesInterface } from "./IEmployee";

export interface RoomInterface {
    ID: number;
    Number: string;
    Name: string;
    TypeID: number;
    Type: TypeInterface;
    BuildingID: number;
    Building: BuildingInterface;
    ServiceDayID: number;
    ServiceDay: ServiceDayInterface;
    PeriodID: number;
    Period: PeriodInterface;
    EmployeeID: number;
    Employee: EmployeesInterface;

}

export interface TypeInterface {
    ID: number;
    Name: string;
    Price: number;
}

export interface BuildingInterface {
    ID: number; 
    Name: string;
}

export interface ServiceDayInterface {
    ID: number;
    Day: string;
}

export interface PeriodInterface {
    ID: number;
    Time: string;
}