import { MedicalRecordInterface } from "./IMedicalRecord";
import { HospitalInterface } from "./IHospital";
import { MedicalHistoryInterface } from "./IMedicalHistory";
import { DoctorInterface } from "./IDoctor";

export interface ReferInterface {

    ReferID: string,

    MedicalRecordID: number,
    MedicalRecord:   MedicalRecordInterface,

    HospitalID: number,
    Hospital :  HospitalInterface,

    MedicalHistoryID: number,
    MedicalHistory: MedicalHistoryInterface,

    DoctorID : number,
    Doctor :  DoctorInterface,

    Date : Date
   
    Cause: string;
   
   }